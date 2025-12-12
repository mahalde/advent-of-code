package datastructures

import (
	"testing"

	"github.com/mahalde/advent-of-code/assert"
)

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList[int]()
	assert.Equals(t, ll.Size, 0)
	assert.Equals(t, ll.Head, (*LinkedListNode[int])(nil))
	assert.Equals(t, ll.Tail, (*LinkedListNode[int])(nil))
}

func TestLinkedList_Append(t *testing.T) {
	ll := NewLinkedList[int]()

	// Append to empty list
	ll.Append(1)
	assert.Equals(t, ll.Size, 1)
	assert.Equals(t, ll.Head.Value, 1)
	assert.Equals(t, ll.Tail.Value, 1)
	assert.Equals(t, ll.Head.Next, (*LinkedListNode[int])(nil))

	// Append second element
	ll.Append(2)
	assert.Equals(t, ll.Size, 2)
	assert.Equals(t, ll.Head.Value, 1)
	assert.Equals(t, ll.Tail.Value, 2)
	assert.Equals(t, ll.Head.Next.Value, 2)
	assert.Equals(t, ll.Tail.Next, (*LinkedListNode[int])(nil))

	// Append third element
	ll.Append(3)
	assert.Equals(t, ll.Size, 3)
	assert.Equals(t, ll.Head.Value, 1)
	assert.Equals(t, ll.Tail.Value, 3)
	assert.Equals(t, ll.Head.Next.Next.Value, 3)
}

func TestLinkedList_Prepend(t *testing.T) {
	ll := NewLinkedList[int]()

	// Prepend to empty list
	ll.Prepend(1)
	assert.Equals(t, ll.Size, 1)
	assert.Equals(t, ll.Head.Value, 1)
	assert.Equals(t, ll.Tail.Value, 1)

	// Prepend second element
	ll.Prepend(2)
	assert.Equals(t, ll.Size, 2)
	assert.Equals(t, ll.Head.Value, 2)
	assert.Equals(t, ll.Tail.Value, 1)
	assert.Equals(t, ll.Head.Next.Value, 1)

	// Prepend third element
	ll.Prepend(3)
	assert.Equals(t, ll.Size, 3)
	assert.Equals(t, ll.Head.Value, 3)
	assert.Equals(t, ll.Tail.Value, 1)
	assert.Equals(t, ll.Head.Next.Next.Value, 1)
}

func TestLinkedList_Find(t *testing.T) {
	ll := NewLinkedList[string]()
	ll.Append("apple")
	ll.Append("banana")
	ll.Append("cherry")

	// Find existing value
	found := ll.Find("banana", func(a, b string) bool { return a == b })
	assert.True(t, found != nil)
	assert.Equals(t, found.Value, "banana")

	// Find non-existing value
	found = ll.Find("date", func(a, b string) bool { return a == b })
	assert.Equals(t, found, (*LinkedListNode[string])(nil))

	// Find in empty list
	emptyLl := NewLinkedList[string]()
	found = emptyLl.Find("anything", func(a, b string) bool { return a == b })
	assert.Equals(t, found, (*LinkedListNode[string])(nil))
}

func TestLinkedList_Remove(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)

	// Remove middle element
	removed := ll.Remove(2, func(a, b int) bool { return a == b })
	assert.True(t, removed)
	assert.Equals(t, ll.Size, 3)
	assert.SlicesEqual(t, ll.ToSlice(), []int{1, 3, 4})

	// Remove head
	removed = ll.Remove(1, func(a, b int) bool { return a == b })
	assert.True(t, removed)
	assert.Equals(t, ll.Size, 2)
	assert.SlicesEqual(t, ll.ToSlice(), []int{3, 4})
	assert.Equals(t, ll.Head.Value, 3)

	// Remove tail
	removed = ll.Remove(4, func(a, b int) bool { return a == b })
	assert.True(t, removed)
	assert.Equals(t, ll.Size, 1)
	assert.SlicesEqual(t, ll.ToSlice(), []int{3})
	assert.Equals(t, ll.Tail.Value, 3)

	// Remove last element
	removed = ll.Remove(3, func(a, b int) bool { return a == b })
	assert.True(t, removed)
	assert.Equals(t, ll.Size, 0)
	assert.Equals(t, ll.Head, (*LinkedListNode[int])(nil))
	assert.Equals(t, ll.Tail, (*LinkedListNode[int])(nil))

	// Try to remove from empty list
	removed = ll.Remove(5, func(a, b int) bool { return a == b })
	assert.False(t, removed)

	// Try to remove non-existing element
	ll.Append(10)
	ll.Append(20)
	removed = ll.Remove(30, func(a, b int) bool { return a == b })
	assert.False(t, removed)
	assert.Equals(t, ll.Size, 2)
}

func TestLinkedList_ToSlice(t *testing.T) {
	ll := NewLinkedList[int]()

	// Empty list
	assert.SlicesEqual(t, ll.ToSlice(), []int{})

	// Single element
	ll.Append(42)
	assert.SlicesEqual(t, ll.ToSlice(), []int{42})

	// Multiple elements
	ll.Append(24)
	ll.Append(12)
	assert.SlicesEqual(t, ll.ToSlice(), []int{42, 24, 12})
}

func TestLinkedList_MixedOperations(t *testing.T) {
	ll := NewLinkedList[int]()

	// Append and prepend
	ll.Append(1)
	ll.Prepend(2)
	ll.Append(3)
	ll.Prepend(4)
	assert.SlicesEqual(t, ll.ToSlice(), []int{4, 2, 1, 3})
	assert.Equals(t, ll.Size, 4)

	// Find and remove
	ll.Remove(2, func(a, b int) bool { return a == b })
	assert.SlicesEqual(t, ll.ToSlice(), []int{4, 1, 3})
	assert.Equals(t, ll.Size, 3)

	// Append more
	ll.Append(5)
	ll.Append(6)
	assert.SlicesEqual(t, ll.ToSlice(), []int{4, 1, 3, 5, 6})

	// Remove head
	ll.Remove(4, func(a, b int) bool { return a == b })
	assert.SlicesEqual(t, ll.ToSlice(), []int{1, 3, 5, 6})

	// Remove tail
	ll.Remove(6, func(a, b int) bool { return a == b })
	assert.SlicesEqual(t, ll.ToSlice(), []int{1, 3, 5})
}
