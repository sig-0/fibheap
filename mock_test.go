package fibheap

type lessDelegate func(Item) bool

// TODO remove
type stringDelegate func() string

// mockItem is a mockable Item
type mockItem struct {
	value int

	lessFn   lessDelegate
	stringFn stringDelegate
}

func (m *mockItem) Less(i Item) bool {
	if m.lessFn != nil {
		return m.lessFn(i)
	}

	other, ok := i.(*mockItem)
	if !ok {
		return false
	}

	return m.value < other.value
}

// TODO remove
func (m *mockItem) String() string {
	if m.stringFn != nil {
		return m.stringFn()
	}

	return ""
}
