package datastructures

type LinkedListNode[T any] struct {
	Value T
	Next  *LinkedListNode[T]
}

type LinkedList[T any] struct {
	Head *LinkedListNode[T]
	Tail *LinkedListNode[T]
	Size int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) Append(value T) {
	newNode := &LinkedListNode[T]{Value: value}
	if ll.Size == 0 {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail.Next = newNode
		ll.Tail = newNode
	}
	ll.Size++
}

func (ll *LinkedList[T]) Prepend(value T) {
	newNode := &LinkedListNode[T]{Value: value}
	if ll.Size == 0 {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		newNode.Next = ll.Head
		ll.Head = newNode
	}
	ll.Size++
}

func (ll *LinkedList[T]) Find(value T, equals func(a, b T) bool) *LinkedListNode[T] {
	current := ll.Head
	for current != nil {
		if equals(current.Value, value) {
			return current
		}
		current = current.Next
	}
	return nil
}

func (ll *LinkedList[T]) Remove(value T, equals func(a, b T) bool) bool {
	if ll.Size == 0 {
		return false
	}

	if equals(ll.Head.Value, value) {
		ll.Head = ll.Head.Next
		ll.Size--
		if ll.Size == 0 {
			ll.Tail = nil
		}
		return true
	}

	current := ll.Head
	for current.Next != nil {
		if equals(current.Next.Value, value) {
			current.Next = current.Next.Next
			if current.Next == nil {
				ll.Tail = current
			}
			ll.Size--
			return true
		}
		current = current.Next
	}

	return false
}

func (ll *LinkedList[T]) ToSlice() []T {
	result := make([]T, 0, ll.Size)
	current := ll.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}
