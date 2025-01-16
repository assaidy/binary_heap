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
	personPriorityQueue.Insert(Person{"person1", 10})
	personPriorityQueue.Insert(Person{"person2", 20})
	personPriorityQueue.Insert(Person{"person3", 5})
	personPriorityQueue.Insert(Person{"person4", 30})

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

func TestHeapDeepVsShallow(t *testing.T) {
	original := []int{10, 20, 5, 30}

	deepHeap := NewHeap(original, func(a, b int) bool { return a > b })
	shallowHeap := NewHeapMute(original, func(a, b int) bool { return a > b })

	shallowHeap.Insert(40)

	if deepHeap.Length() == shallowHeap.Length() {
		t.Errorf("Deep heap was unexpectedly affected by shallow heap changes")
	}

	var deepResult, shallowResult []int
	for !deepHeap.IsEmpty() {
		deepResult = append(deepResult, deepHeap.Pop())
	}
	for !shallowHeap.IsEmpty() {
		shallowResult = append(shallowResult, shallowHeap.Pop())
	}

	expectedDeep := []int{30, 20, 10, 5}
	if !reflect.DeepEqual(deepResult, expectedDeep) {
		t.Errorf("Deep heap failed: got %v, want %v", deepResult, expectedDeep)
	}

	expectedShallow := []int{40, 30, 20, 10, 5}
	if !reflect.DeepEqual(shallowResult, expectedShallow) {
		t.Errorf("Shallow heap failed: got %v, want %v", shallowResult, expectedShallow)
	}
}
