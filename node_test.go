package fibheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_New(t *testing.T) {
	t.Parallel()

	item := &mockItem{}
	n := newNode(item)

	assert.Equal(t, item, n.item)
	assert.Equal(t, n, n.right)
	assert.Equal(t, n, n.left)
	assert.Nil(t, n.parent)
}

func TestNode_RemoveFromSiblings(t *testing.T) {
	t.Parallel()

	item := &mockItem{}
	n := newNode(item)
	left := newNode(item)
	right := newNode(item)

	// Create initial sibling ties
	n.left = left
	n.right = right

	n.left.left = n.right
	n.right.right = n.left

	// Remove the current node from siblings
	n.removeFromSiblings()

	// Make sure the node is not referenced
	// by the sibling nodes
	assert.Equal(t, right, left.right)
	assert.Equal(t, right, left.left)
	assert.Equal(t, left, right.left)
	assert.Equal(t, right, right.right)

	// Make sure the node still has its original
	// sibling ties
	assert.Equal(t, left, n.left)
	assert.Equal(t, right, n.right)
}

func TestNode_RemoveSiblings(t *testing.T) {
	t.Parallel()

	item := &mockItem{}
	n := newNode(item)
	left := newNode(item)
	right := newNode(item)

	// Create initial sibling ties
	n.left = left
	n.right = right

	// Remove sibling ties
	n.removeSiblings()

	// Make sure the node does not have sibling ties
	assert.Equal(t, n, n.left)
	assert.Equal(t, n, n.right)
}

func TestNode_AddSibling(t *testing.T) {
	t.Parallel()

	item := &mockItem{}

	t.Run("invalid sibling", func(t *testing.T) {
		t.Parallel()

		n := newNode(item)

		// Add invalid sibling
		n.addSibling(nil)

		assert.Equal(t, n, n.left)
		assert.Equal(t, n, n.right)
	})

	t.Run("no initial siblings", func(t *testing.T) {
		t.Parallel()

		n := newNode(item)
		sibling := newNode(item)

		// Add sibling
		n.addSibling(sibling)

		assert.Equal(t, sibling, n.left)
		assert.Equal(t, sibling, n.right)

		assert.Equal(t, n, sibling.left)
		assert.Equal(t, n, sibling.right)
	})

	t.Run("node has initial sibling", func(t *testing.T) {
		t.Parallel()

		getSiblings := func(count int) []*node {
			siblings := make([]*node, count)

			for i := 0; i < count; i++ {
				siblings[i] = newNode(item)
			}

			return siblings
		}

		n := newNode(item)
		siblings := getSiblings(10)

		// Add siblings
		for _, sibling := range siblings {
			n.addSibling(sibling)
		}

		// Make sure all siblings are there
		siblingCount := 0
		currentNode := n.right
		for currentNode != n {
			siblingCount++

			currentNode = currentNode.right
		}

		assert.Equal(t, len(siblings), siblingCount)
	})
}

func TestNode_AddChild(t *testing.T) {
	t.Parallel()

	item := &mockItem{}

	t.Run("invalid child", func(t *testing.T) {
		t.Parallel()

		n := newNode(item)

		// Add empty child
		n.addChild(nil)

		assert.Equal(t, uint(0), n.degree)
		assert.Nil(t, n.child)
	})

	t.Run("no initial children", func(t *testing.T) {
		t.Parallel()

		n := newNode(item)
		refNode := newNode(item)

		assert.Nil(t, refNode.child)

		// Add the node as a child
		refNode.addChild(n)

		assert.Equal(t, n, refNode.child)
		assert.Equal(t, refNode, n.parent)
		assert.Equal(t, uint(1), refNode.degree)

		assert.Equal(t, n, n.left)
		assert.Equal(t, n, n.right)
	})

	t.Run("parent has initial children", func(t *testing.T) {
		t.Parallel()

		n := newNode(item)
		refNode := newNode(item)
		child := newNode(item)

		// Add an initial child
		refNode.addChild(child)
		assert.NotNil(t, refNode.child)

		// Add the node as a child
		refNode.addChild(n)

		assert.Equal(t, n, child.left)
		assert.Equal(t, n, child.right)

		assert.Equal(t, child, n.left)
		assert.Equal(t, child, n.right)

		assert.Equal(t, refNode, n.parent)
		assert.Equal(t, uint(2), refNode.degree)
	})
}
