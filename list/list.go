package list

import ()

// List holds the internal structure of a doubly-linked list
type List struct {
	head *node
	tail *node
}

// Internal struct as a node
type node struct {
	data interface{}
	next *node
	prev *node
}

// Initialize a empty List
func New() *List {
	return &List{head: nil, tail: nil}
}

// Clone the content of an existing list.  Shallow copy implemented.
func (list *List) Clone() *List {
	newList := New()

	for e := range list.Iter() {
		newList.Push(e)
	}

	return newList
}

// Push data into the tail of the list
func (list *List) Push(data interface{}) {
	node := node{data: data, next: nil, prev: nil}
	if list.head == nil {
		list.head = &node
		list.tail = &node
	} else {
		list.tail.next = &node
		node.prev = list.tail
		list.tail = &node
	}
}

// Pop data from the tail of the list.  Return error if empty list.
func (list *List) Pop() (interface{}, error) {
	if list.head == nil {
		return nil, emptyList()
	}

	if list.head == list.tail {
		data := list.head.data
		list.head = nil
		list.tail = nil
		return data, nil
	} else {
		current := list.tail
		data := current.data
		list.tail = current.prev
		list.tail.next = nil
		current = nil
		return data, nil
	}
}

// Unshift data into the head of the list
func (list *List) Unshift(data interface{}) {
	node := node{data: data, prev: nil, next: nil}
	if list.head == nil {
		list.head = &node
		list.tail = &node
	} else {
		list.head.prev = &node
		node.next = list.head
		list.head = &node
	}
}

// Shift data from the head of the list.  Return error if empty list.
func (list *List) Shift() (interface{}, error) {
	if list.head == nil {
		return nil, emptyList()
	}

	if list.head == list.tail {
		data := list.head.data
		list.head = nil
		list.tail = nil
		return data, nil
	} else {
		current := list.head
		data := current.data
		list.head = current.next
		list.head.prev = nil
		current = nil
		return data, nil
	}
}

// Get the length of the list.
func (list *List) Len() uint {
	var count uint = 0
	current := list.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

// Check whether the list is empty.
func (list *List) IsEmpty() bool {
	return list.head == nil
}

// Iterate through the list and get the elements.
func (list *List) Iter() <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		current := list.head
		for current != nil {
			ch <- current.data
			current = current.next
		}
		close(ch)
	}()

	return ch
}

// An alias of Iter
func (list *List) Each() <-chan interface{} {
	return list.Iter()
}

// Get the element at ``index``.  Returns error if ``index`` out of range
func (list *List) At(index uint) (interface{}, error) {
	if list.isIndexOutOfRange(index) {
		return nil, indexOutOfRange()
	}

	var count uint = 0
	current := list.head
	for count < index {
		current = current.next
		count++
	}

	return current.data, nil
}

// Remove the element at ``index`` from the list.  Returns errors if ``index``
// out of range.
func (list *List) RemoveAt(index uint) (interface{}, error) {
	if list.isIndexOutOfRange(index) {
		return nil, indexOutOfRange()
	}

	var previous *node = nil
	var current *node = list.head
	var count uint = 0
	for count < index {
		previous = current
		current = current.next
		count++
	}

	if current == list.head { // index == 0
		return list.Shift()
	} else if current == list.tail { // index == Len() - 1
		return list.Pop()
	} else {
		data := current.data
		previous.next = current.next
		current.next.prev = previous
		current.prev = nil
		current.next = nil
		current = nil
		return data, nil
	}
}

func (list *List) isIndexOutOfRange(index uint) bool {
	if list.Len() == 0 {
		return true
	}

	if index > list.Len()-1 {
		return true
	}

	return false
}
