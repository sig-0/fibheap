package fibheap

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// generateMockItems generates mock items with random int values
func generateMockItems(count int) []*mockItem {
	items := make([]*mockItem, count)

	for i := 0; i < count; i++ {
		items[i] = &mockItem{
			value: rand.Intn(count),
		}
	}

	return items
}

func TestHeap_Push(t *testing.T) {
	t.Parallel()

	numItems := 100
	items := generateMockItems(numItems)

	h := NewHeap()

	for _, item := range items {
		h.Push(item)
	}

	assert.Equal(t, uint(numItems), h.size)
}

func TestHeap_Pop(t *testing.T) {
	t.Parallel()

	item := &mockItem{}
	h := NewHeap()

	h.Push(item)

	popped := h.Pop()

	assert.Equal(t, item, popped)
	assert.Equal(t, uint(0), h.size)
	assert.Nil(t, h.Pop())
}

func TestHeap_Peek(t *testing.T) {
	t.Parallel()

	t.Run("empty heap", func(t *testing.T) {
		t.Parallel()

		h := NewHeap()

		assert.Equal(t, uint(0), h.Size())
		assert.Nil(t, h.Peek())
	})

	t.Run("min heap", func(t *testing.T) {
		t.Parallel()

		items := generateMockItems(10)

		h := NewHeap()

		// Find min and add items
		min := items[0]
		for _, item := range items {
			if item.value < min.value {
				min = item
			}

			h.Push(item)
		}

		assert.Equal(t, uint(len(items)), h.Size())

		assert.Equal(t, min, h.Peek())
	})

	t.Run("max heap", func(t *testing.T) {
		t.Parallel()

		items := generateMockItems(10)

		h := NewHeap()

		// Find min and add items
		max := items[0]
		for _, item := range items {
			item := item

			// Set it to be a max heap
			item.lessFn = func(otherRaw Item) bool {
				other, _ := otherRaw.(*mockItem)

				return item.value >= other.value
			}

			if item.value > max.value {
				max = item
			}

			h.Push(item)
		}

		assert.Equal(t, uint(len(items)), h.Size())

		assert.Equal(t, max.value, h.Peek().(*mockItem).value)
	})
}

func TestHeap_Merge(t *testing.T) {
	t.Parallel()

	items := []*mockItem{
		{
			value: 10,
		},
		{
			value: 5,
		},
	}

	heaps := []*Heap{
		{
			entry: newNode(items[0]),
			size:  1,
		},
		{
			entry: newNode(items[1]),
			size:  1,
		},
	}

	// Merge the heaps
	heaps[0].Merge(heaps[1])

	// Make sure the heap size increased
	heap := heaps[0]

	assert.Equal(t, uint(2), heap.Size())
	assert.Equal(t, items[1], heap.Pop())
}

func TestHeap_Merge_Empty(t *testing.T) {
	t.Parallel()

	item := &mockItem{}
	h := NewHeap()

	h.Push(item)

	assertCommon := func() {
		assert.Equal(t, uint(1), h.Size())
		assert.Equal(t, item, h.Peek())
	}

	assertCommon()

	// Empty heap
	h.Merge(NewHeap())
	assertCommon()

	// Invalid heap
	h.Merge(nil)
	assertCommon()
}

func TestHeap_Merge_AssignEntry(t *testing.T) {
	t.Parallel()

	item := &mockItem{}
	h := NewHeap()

	heapToMerge := NewHeap()
	heapToMerge.Push(item)

	assert.Equal(t, uint(1), heapToMerge.Size())

	// Merge the heaps
	h.Merge(heapToMerge)

	// Make sure the heap size increased
	assert.Equal(t, uint(1), h.Size())
	assert.Equal(t, item, h.Pop())
}

func TestHeap_PushPop(t *testing.T) {
	t.Parallel()

	t.Run("min heap", func(t *testing.T) {
		t.Parallel()

		h := NewHeap()

		items := generateMockItems(1000)

		minSorted := make([]*mockItem, len(items))
		copy(minSorted, items)

		// Sort the array ascending
		sort.SliceStable(minSorted, func(i, j int) bool {
			return minSorted[i].value < minSorted[j].value
		})

		for _, item := range items {
			h.Push(item)
		}

		// Pop them off one by one
		for _, item := range minSorted {
			popped := h.Pop().(*mockItem)
			assert.Equal(t, item.value, popped.value)
		}

		assert.Equal(t, uint(0), h.Size())
	})

	t.Run("max heap", func(t *testing.T) {
		t.Parallel()

		items := generateMockItems(1000)

		h := NewHeap()

		maxSorted := make([]*mockItem, len(items))
		copy(maxSorted, items)

		// Sort the array ascending
		sort.SliceStable(maxSorted, func(i, j int) bool {
			return maxSorted[i].value >= maxSorted[j].value
		})

		for _, item := range items {
			item := item

			// Set it to be a max heap
			item.lessFn = func(otherRaw Item) bool {
				other, _ := otherRaw.(*mockItem)

				return item.value >= other.value
			}

			h.Push(item)
		}

		// Pop them off one by one
		for _, item := range maxSorted {
			popped := h.Pop().(*mockItem)
			assert.Equal(t, item.value, popped.value)
		}

		assert.Equal(t, uint(0), h.Size())
	})
}

func TestHeap_Clear(t *testing.T) {
	t.Parallel()

	items := generateMockItems(1000)

	h := NewHeap()

	for _, item := range items {
		h.Push(item)
	}

	assert.Equal(t, uint(len(items)), h.Size())

	h.Clear()

	assert.Equal(t, uint(0), h.Size())
	assert.Nil(t, h.entry)
}
