package fibheap

// node represents a single node in a circular linked list.
// It has references to its parent (if any), and left-most child (if any)
type node struct {
	item   Item
	degree uint

	child  *node // left-most child
	parent *node

	right *node
	left  *node
}

// newNode creates a new unconnected node
func newNode(item Item) *node {
	n := &node{
		item: item,
	}

	n.left = n
	n.right = n

	return n
}

// removeFromSiblings removes the current node's reference
// from its siblings. Note that this method preserves
// the node's original sibling ties (left, right)
func (n *node) removeFromSiblings() {
	if n.right == n {
		// Node does not have siblings
		return
	}

	// Fix sibling links
	n.left.right = n.right
	n.right.left = n.left
}

// removeSiblings removes the current node's sibling references
// and makes the node as its only sibling
func (n *node) removeSiblings() {
	if n.right == n {
		// Node does not have siblings
		return
	}

	n.left, n.right = n, n
}

// addChild adds a new child to the current node
func (n *node) addChild(child *node) {
	if child == nil {
		return
	}

	// Increase the parent degree
	n.degree++

	child.parent = n

	if n.child == nil {
		// Parent does not have children,
		// so this is the only child
		n.child = child
		child.left, child.right = child, child

		return
	}

	// Parent has existing children, add the node as their sibling
	n.child.addSibling(child)
}

// addSibling adds a node as a sibling node
func (n *node) addSibling(sibling *node) {
	if sibling == nil {
		return
	}

	// n.left -> sibling -> n
	n.left.right = sibling
	sibling.right = n

	// n.left -> <- sibling -> <- n
	sibling.left = n.left
	n.left = sibling
}
