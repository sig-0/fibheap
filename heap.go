package fibheap

import "math"

// Heap is an instance of the Fibonacci heap
type Heap struct {
	size  uint  // number of nodes
	entry *node // min or max node
}

// NewHeap creates a new Fibonacci heap instance
func NewHeap() *Heap {
	return &Heap{
		size:  0,
		entry: nil,
	}
}

// Size returns the number of nodes in the heap
func (h *Heap) Size() uint {
	return h.size
}

// Push adds a new item to the heap
func (h *Heap) Push(item Item) {
	// Create a new node for the item
	n := newNode(item)

	h.size++

	if h.entry == nil {
		// Heap is empty, and this is the first element
		h.entry = n

		return
	}

	// Add the new node to the top-level root list
	h.entry.addSibling(n)

	// Attempt to swap the entry node
	h.attemptToSwapEntry(n)
}

// attemptToSwapEntry attempts to swap the entry node with the given element
func (h *Heap) attemptToSwapEntry(n *node) {
	if n.item.Less(h.entry.item) {
		h.entry = n
	}
}

// Peek returns the lowest / highest value item in the heap,
// depending on the configuration, but does not remove it
func (h *Heap) Peek() Item {
	if h.entry != nil {
		return h.entry.item
	}

	return nil
}

// Pop removes the lowest / highest value item in the heap
// depending on the configuration
func (h *Heap) Pop() Item {
	// Check if there is an entry element at all
	if h.entry == nil {
		return nil
	}

	// 1. Save a reference the entry item
	item := h.entry.item

	// 2. Append the children to the top-level root list
	extractNode := func(n *node) {
		// Add the node to the root list
		h.entry.addSibling(n)

		// Remove the parent
		n.parent = nil
	}

	currentChild := h.entry.child
	if currentChild != nil {
		lastChild := h.entry.child.left

		for currentChild != lastChild {
			// Save the reference to the next node
			next := currentChild.right

			// Extract the subtree
			extractNode(currentChild)

			// Move over to the next sibling
			currentChild = next
		}

		// Apply the processing on the last node
		extractNode(currentChild)
	}

	// Remove the entry root node from the root list
	h.entry.removeFromSiblings()

	// 3. Perform heap-tidy
	if h.entry == h.entry.right {
		// The entry node was the last node
		h.entry = nil
	} else {
		// Make the next node a temporary entry
		h.entry = h.entry.right

		// Consolidate the heap
		h.consolidate()
	}

	// Decrease the heap size
	h.size--

	return item
}

// consolidate consolidates the heap by simplifying and merging
func (h *Heap) consolidate() {
	// The max degree is always bound by logN
	maxDegree := int(math.Floor(math.Log(float64(h.size)) / math.Log((1+math.Sqrt(5))/2)))

	degreeMap := make([]*node, maxDegree+2)

	mergeTrees := func(currentNode *node) {
		currentNode.removeFromSiblings()
		currentNode.removeSiblings()

		degree := currentNode.degree

		for degreeMap[degree] != nil {
			otherNode := degreeMap[degree]

			if otherNode.item.Less(currentNode.item) {
				// Swap the pointers
				otherNode, currentNode = currentNode, otherNode
			}

			h.linkNodes(currentNode, otherNode)

			degreeMap[degree] = nil
			degree++
		}

		degreeMap[currentNode.degree] = currentNode
	}

	currentNode := h.entry
	lastNode := h.entry.left

	for currentNode != lastNode {
		next := currentNode.right

		mergeTrees(currentNode)

		currentNode = next
	}

	// Apply for the last node
	mergeTrees(currentNode)

	// Find the new entry element
	h.entry = nil

	for _, n := range degreeMap {
		if n == nil {
			continue
		}

		if h.entry == nil {
			h.entry = n

			continue
		}

		h.entry.addSibling(n)

		if n.item.Less(h.entry.item) {
			h.entry = n
		}
	}
}

// linkNodes creates a parent-child relationship between
// given nodes, removing the child
// from the top-level root list
func (h *Heap) linkNodes(parent, child *node) {
	// Remove all sibling ties
	child.removeFromSiblings()
	child.removeSiblings()

	// Add the child node to the new parent
	parent.addChild(child)
}

// Merge combines the two heaps
func (h *Heap) Merge(heap *Heap) {
	if heap == nil || heap.entry == nil {
		// No need to merge heaps
		// if one is empty
		return
	}

	if h.entry == nil {
		h.entry = heap.entry
		h.size = heap.size

		return
	}

	// Try to find the new entry node
	oldEntry := h.entry
	h.attemptToSwapEntry(heap.entry)

	// Join the root lists of the heaps
	last := oldEntry.left
	heap.entry.left = oldEntry.left

	oldEntry.left.right = heap.entry
	oldEntry.left = last
	oldEntry.left.right = oldEntry

	// Update the total node count
	h.size += heap.size
}

// Clear clears the heap of any elements
func (h *Heap) Clear() {
	h.entry = nil
	h.size = 0
}
