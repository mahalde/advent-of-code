package datastructures

import (
	"testing"

	"github.com/mahalde/advent-of-code/assert"
)

func TestNewDoublyLinkedList(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	assert.Equals(t, dll.Size, 0)
	assert.Equals(t, dll.Head, (*DoublyLinkedListNode[int])(nil))
	assert.Equals(t, dll.Tail, (*DoublyLinkedListNode[int])(nil))
}

func TestDoublyLinkedList_Append(t *testing.T) {
	dll := NewDoublyLinkedList[int]()

	// Append to empty list
	dll.Append(1)
	assert.Equals(t, dll.Size, 1)
	assert.Equals(t, dll.Head.Value, 1)
	assert.Equals(t, dll.Tail.Value, 1)
	assert.Equals(t, dll.Head.Previous, (*DoublyLinkedListNode[int])(nil))
	assert.Equals(t, dll.Head.Next, (*DoublyLinkedListNode[int])(nil))

	// Append second element
	dll.Append(2)
	assert.Equals(t, dll.Size, 2)
	assert.Equals(t, dll.Head.Value, 1)
	assert.Equals(t, dll.Tail.Value, 2)
	assert.Equals(t, dll.Head.Next.Value, 2)
	assert.Equals(t, dll.Tail.Previous.Value, 1)

	// Append third element
	dll.Append(3)
	assert.Equals(t, dll.Size, 3)
	assert.Equals(t, dll.Head.Value, 1)
	assert.Equals(t, dll.Tail.Value, 3)
	assert.Equals(t, dll.Head.Next.Next.Value, 3)
	assert.Equals(t, dll.Tail.Previous.Previous.Value, 1)
}

func TestDoublyLinkedList_Prepend(t *testing.T) {
	dll := NewDoublyLinkedList[int]()

	// Prepend to empty list
	dll.Prepend(1)
	assert.Equals(t, dll.Size, 1)
	assert.Equals(t, dll.Head.Value, 1)
	assert.Equals(t, dll.Tail.Value, 1)

	// Prepend second element
	dll.Prepend(2)
	assert.Equals(t, dll.Size, 2)
	assert.Equals(t, dll.Head.Value, 2)
	assert.Equals(t, dll.Tail.Value, 1)
	assert.Equals(t, dll.Head.Next.Value, 1)
	assert.Equals(t, dll.Tail.Previous.Value, 2)

	// Prepend third element
	dll.Prepend(3)
	assert.Equals(t, dll.Size, 3)
	assert.Equals(t, dll.Head.Value, 3)
	assert.Equals(t, dll.Tail.Value, 1)
	assert.Equals(t, dll.Head.Next.Next.Value, 1)
	assert.Equals(t, dll.Tail.Previous.Previous.Value, 3)
}

func TestDoublyLinkedList_Remove(t *testing.T) {
	dll := NewDoublyLinkedList[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)
	dll.Append(4)

	// Remove middle node
	nodeToRemove := dll.Head.Next // value 2
	dll.Remove(nodeToRemove)
	assert.Equals(t, dll.Size, 3)
	assert.SlicesEqual(t, dll.ToSlice(), []int{1, 3, 4})
	assert.Equals(t, dll.Head.Value, 1)
	assert.Equals(t, dll.Tail.Value, 4)
	assert.Equals(t, dll.Head.Next.Value, 3)
	assert.Equals(t, dll.Tail.Previous.Value, 3)

	// Remove head
	dll.Remove(dll.Head)
	assert.Equals(t, dll.Size, 2)
	assert.SlicesEqual(t, dll.ToSlice(), []int{3, 4})
	assert.Equals(t, dll.Head.Value, 3)
	assert.Equals(t, dll.Tail.Value, 4)

	// Remove tail
	dll.Remove(dll.Tail)
	assert.Equals(t, dll.Size, 1)
	assert.SlicesEqual(t, dll.ToSlice(), []int{3})
	assert.Equals(t, dll.Head.Value, 3)
	assert.Equals(t, dll.Tail.Value, 3)

	// Remove last element
	dll.Remove(dll.Head)
	assert.Equals(t, dll.Size, 0)
	assert.Equals(t, dll.Head, (*DoublyLinkedListNode[int])(nil))
	assert.Equals(t, dll.Tail, (*DoublyLinkedListNode[int])(nil))
}

func TestDoublyLinkedList_Find(t *testing.T) {
	dll := NewDoublyLinkedList[string]()
	dll.Append("apple")
	dll.Append("banana")
	dll.Append("cherry")

	// Find existing value
	found := dll.Find("banana", func(a, b string) bool { return a == b })
	assert.True(t, found != nil)
	assert.Equals(t, found.Value, "banana")

	// Find non-existing value
	found = dll.Find("date", func(a, b string) bool { return a == b })
	assert.Equals(t, found, (*DoublyLinkedListNode[string])(nil))

	// Find in empty list
	emptyDll := NewDoublyLinkedList[string]()
	found = emptyDll.Find("anything", func(a, b string) bool { return a == b })
	assert.Equals(t, found, (*DoublyLinkedListNode[string])(nil))
}

func TestDoublyLinkedList_ToSlice(t *testing.T) {
	dll := NewDoublyLinkedList[int]()

	// Empty list
	assert.SlicesEqual(t, dll.ToSlice(), []int{})

	// Single element
	dll.Append(42)
	assert.SlicesEqual(t, dll.ToSlice(), []int{42})

	// Multiple elements
	dll.Append(24)
	dll.Append(12)
	assert.SlicesEqual(t, dll.ToSlice(), []int{42, 24, 12})
}

func TestDoublyLinkedList_MixedOperations(t *testing.T) {
	dll := NewDoublyLinkedList[int]()

	// Append and prepend
	dll.Append(1)
	dll.Prepend(2)
	dll.Append(3)
	dll.Prepend(4)
	assert.SlicesEqual(t, dll.ToSlice(), []int{4, 2, 1, 3})
	assert.Equals(t, dll.Size, 4)

	// Find and remove
	node := dll.Find(2, func(a, b int) bool { return a == b })
	dll.Remove(node)
	assert.SlicesEqual(t, dll.ToSlice(), []int{4, 1, 3})
	assert.Equals(t, dll.Size, 3)

	// Append more
	dll.Append(5)
	dll.Append(6)
	assert.SlicesEqual(t, dll.ToSlice(), []int{4, 1, 3, 5, 6})

	// Remove head
	dll.Remove(dll.Head)
	assert.SlicesEqual(t, dll.ToSlice(), []int{1, 3, 5, 6})

	// Remove tail
	dll.Remove(dll.Tail)
	assert.SlicesEqual(t, dll.ToSlice(), []int{1, 3, 5})
}
