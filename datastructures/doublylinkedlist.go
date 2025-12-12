package datastructures

type DoublyLinkedListNode[T any] struct {
	Value    T
	Previous *DoublyLinkedListNode[T]
	Next     *DoublyLinkedListNode[T]
}

type DoublyLinkedList[T any] struct {
	Head *DoublyLinkedListNode[T]
	Tail *DoublyLinkedListNode[T]
	Size int
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (dll *DoublyLinkedList[T]) Append(value T) {
	newNode := &DoublyLinkedListNode[T]{Value: value}
	if dll.Size == 0 {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Previous = dll.Tail
		dll.Tail.Next = newNode
		dll.Tail = newNode
	}
	dll.Size++
}

func (dll *DoublyLinkedList[T]) Prepend(value T) {
	newNode := &DoublyLinkedListNode[T]{Value: value}
	if dll.Size == 0 {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Next = dll.Head
		dll.Head.Previous = newNode
		dll.Head = newNode
	}
	dll.Size++
}

func (dll *DoublyLinkedList[T]) Remove(node *DoublyLinkedListNode[T]) {
	if node.Previous != nil {
		node.Previous.Next = node.Next
	} else {
		dll.Head = node.Next
	}
	if node.Next != nil {
		node.Next.Previous = node.Previous
	} else {
		dll.Tail = node.Previous
	}
	dll.Size--
}

func (dll *DoublyLinkedList[T]) Find(value T, equals func(a, b T) bool) *DoublyLinkedListNode[T] {
	current := dll.Head
	for current != nil {
		if equals(current.Value, value) {
			return current
		}
		current = current.Next
	}
	return nil
}

func (dll *DoublyLinkedList[T]) ToSlice() []T {
	slice := make([]T, 0, dll.Size)
	current := dll.Head
	for current != nil {
		slice = append(slice, current.Value)
		current = current.Next
	}
	return slice
}
