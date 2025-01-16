package binary_heap

import (
	"reflect"
	"testing"
)

type Person struct {
	name string
	age  int
}

func TestMaxHeap(t *testing.T) {
	maxHeap := NewHeap([]int{10, 20, 5, 30}, func(a, b int) bool { return a > b })
	var result []int
	for !maxHeap.IsEmpty() {
		result = append(result, maxHeap.Pop())
	}
	expected := []int{30, 20, 10, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MaxHeap failed: got %v, want %v", result, expected)
	}
}

func TestMinHeap(t *testing.T) {
	minHeap := NewHeap([]int{10, 20, 5, 30}, func(a, b int) bool { return a < b })
	var result []int
	for !minHeap.IsEmpty() {
		result = append(result, minHeap.Pop())
	}
	expected := []int{5, 10, 20, 30}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MinHeap failed: got %v, want %v", result, expected)
	}
}

func TestHeapSort(t *testing.T) {
	ints := []int{10, 20, 5, 30}
	HeapSort(ints, func(a, b int) bool { return a < b })
	expected := []int{5, 10, 20, 30}
	if !reflect.DeepEqual(ints, expected) {
		t.Errorf("HeapSort failed: got %v, want %v", ints, expected)
	}
}

func TestPersonPriorityQueue(t *testing.T) {
	personPriorityQueue := NewHeap([]Person{}, func(a, b Person) bool { return a.age > b.age })
	personPriorityQueue.Push(Person{"person1", 10})
	personPriorityQueue.Push(Person{"person2", 20})
	personPriorityQueue.Push(Person{"person3", 5})
	personPriorityQueue.Push(Person{"person4", 30})

	var result []Person
	for !personPriorityQueue.IsEmpty() {
		result = append(result, personPriorityQueue.Pop())
	}

	expected := []Person{
		{"person4", 30},
		{"person2", 20},
		{"person1", 10},
		{"person3", 5},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PersonPriorityQueue failed: got %v, want %v", result, expected)
	}
}

func TestEmptyHeap(t *testing.T) {
	emptyHeap := NewHeap([]int{}, func(a, b int) bool { return a > b })
	if !emptyHeap.IsEmpty() {
		t.Errorf("EmptyHeap failed: expected empty heap but got non-empty")
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic on Pop from empty heap, but no panic occurred")
		}
	}()
	emptyHeap.Pop()
}
