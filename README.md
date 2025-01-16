# Binary Heap Implementation

This package provides a generic implementation of a binary heap in Go, supporting both max-heaps and min-heaps, as well as heap sort functionality. The implementation is versatile and can handle custom types using user-defined comparison functions.

## Features

- Max-Heap and Min-Heap support
- Generic implementation for custom data types
- Heap sort for in-place sorting
- Two initialization options:
  - **Deep copy heap** (`NewHeap`): Creates a heap with a new copy of the input slice.
  - **Shallow copy heap** (`NewHeapMute`): Directly modifies the input slice.

---

## Installation

```bash
go get github.com/assaidy/binary_heap
```

---

## Usage Example

### Max-Heap Example

```go
package main

import (
	"fmt"
	"github.com/assaidy/binary_heap"
)

func main() {
	maxHeap := binary_heap.NewHeap([]int{10, 20, 5, 30}, func(a, b int) bool { return a > b })
	for !maxHeap.IsEmpty() {
		fmt.Println(maxHeap.Pop()) // Output: 30, 20, 10, 5
	}
}
```

### Min-Heap Example

```go
package main

import (
	"fmt"
	"github.com/assaidy/binary_heap"
)

func main() {
	minHeap := binary_heap.NewHeap([]int{10, 20, 5, 30}, func(a, b int) bool { return a < b })
	for !minHeap.IsEmpty() {
		fmt.Println(minHeap.Pop()) // Output: 5, 10, 20, 30
	}
}
```

### Custom Data Type with Priority Queue

```go
package main

import (
	"fmt"
	"github.com/assaidy/binary_heap"
)

type Person struct {
	name string
	age  int
}

func main() {
	personHeap := binary_heap.NewHeap([]Person{}, func(a, b Person) bool { return a.age > b.age })
	personHeap.Insert(Person{"Alice", 30})
	personHeap.Insert(Person{"Bob", 25})
	personHeap.Insert(Person{"Charlie", 35})

	for !personHeap.IsEmpty() {
		fmt.Println(personHeap.Pop()) // Output: Charlie, Alice, Bob
	}
}
```

### Heap Sort Example

```go
package main

import (
	"fmt"
	"github.com/assaidy/binary_heap"
)

func main() {
	ints := []int{10, 20, 5, 30}
	binary_heap.HeapSort(ints, func(a, b int) bool { return a < b })
	fmt.Println(ints) // Output: [5, 10, 20, 30]
}
```

---

## Testing

The package includes a comprehensive test suite. To run the tests:

```bash
go test ./...
```

