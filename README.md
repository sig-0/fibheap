## Overview

[![codecov](https://codecov.io/gh/madz-lab/fibheap/branch/main/graph/badge.svg?token=U85BAYBPPJ)](https://codecov.io/gh/madz-lab/fibheap)

`fibheap` is small and simple [Fibonacci Heap](https://en.wikipedia.org/wiki/Fibonacci_heap) implementation, written in
Go.

It can be utilized as a min or max heap, depending on the implementation of the `Item.Less` method.

Fibonacci heaps are a type of heap data structure that provide faster insertion and deletion operations compared to
binary heaps, but at the cost of increased space complexity. Depending on the concrete implementation of the binary
heap,
it is possible to achieve better results compared to Fibonacci heaps (ex. `container/heap`), because of underlying array
optimizations.

| Operation  | Time Complexity (worst case) |
|------------|------------------------------|
| Insert     | O(1)                         |
| Find min   | O(1)                         |
| Delete min | O(logN)                      |
| Merge      | O(1)                         |

### Internal Representation

The heap itself is internally represented as a doubly linked circular list of nodes.
Parent nodes have references to the left-most child. Each child has a reference to their parent.

An example heap representation after adding elements `[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]`, and removing the lowest
element (`0`):
![banner](.github/banner.png)

### Benchmarks

```bash
goos: darwin
goarch: amd64
pkg: github.com/madz-lab/fibheap
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkHeap_Push1000-8           25191             48070 ns/op
BenchmarkHeap_Push10000-8           2234            521024 ns/op
BenchmarkHeap_Push100000-8           178           6716704 ns/op
BenchmarkHeap_Push1000000-8           18          60324315 ns/op
BenchmarkHeap_Pop1000-8             2694            428768 ns/op
BenchmarkHeap_Pop10000-8             189           6308310 ns/op
BenchmarkHeap_Pop100000-8              9         115665221 ns/op
BenchmarkHeap_Pop1000000-8             1        2242760745 ns/op
PASS
ok      github.com/madz-lab/fibheap     14.434s
```