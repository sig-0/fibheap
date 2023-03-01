package fibheap

import "testing"

func benchmarkPush(b *testing.B, items []*mockItem) {
	b.Helper()

	for i := 0; i < b.N; i++ {
		h := NewHeap()
		for _, item := range items {
			h.Push(item)
		}
	}
}

func BenchmarkHeap_Push1000(b *testing.B) {
	items := generateMockItems(1000)

	b.ResetTimer()
	benchmarkPush(b, items)
}

func BenchmarkHeap_Push10000(b *testing.B) {
	items := generateMockItems(10000)

	b.ResetTimer()
	benchmarkPush(b, items)
}

func BenchmarkHeap_Push100000(b *testing.B) {
	items := generateMockItems(100000)

	b.ResetTimer()
	benchmarkPush(b, items)
}

func BenchmarkHeap_Push1000000(b *testing.B) {
	items := generateMockItems(1000000)

	b.ResetTimer()
	benchmarkPush(b, items)
}

func benchmarkPop(b *testing.B, items []*mockItem) {
	b.Helper()

	for i := 0; i < b.N; i++ {
		h := NewHeap()

		b.StopTimer()

		for _, item := range items {
			h.Push(item)
		}

		b.StartTimer()

		for h.Size() > 0 {
			h.Pop()
		}
	}
}

func BenchmarkHeap_Pop1000(b *testing.B) {
	items := generateMockItems(1000)

	b.ResetTimer()
	benchmarkPop(b, items)
}

func BenchmarkHeap_Pop10000(b *testing.B) {
	items := generateMockItems(10000)

	b.ResetTimer()
	benchmarkPop(b, items)
}

func BenchmarkHeap_Pop100000(b *testing.B) {
	items := generateMockItems(100000)

	b.ResetTimer()
	benchmarkPop(b, items)
}

func BenchmarkHeap_Pop1000000(b *testing.B) {
	items := generateMockItems(1000000)

	b.ResetTimer()
	benchmarkPop(b, items)
}
