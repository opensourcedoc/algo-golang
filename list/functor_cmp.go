package list

// Check whether two lists are equal.  If both the lengths and data of the
// two list are identical, the two lists are equal.
// Implement cmp function object to use this method.
func (list *List) Equal(other *List, cmp func(interface{}, interface{}) (int, error)) (bool, error) {
	lPtr := list.head
	rPtr := other.head

	for lPtr != nil && rPtr != nil {
		b, err := cmp(lPtr.data, rPtr.data)
		if err != nil {
			return false, err
		}

		if b != 0 {
			return false, nil
		}

		lPtr = lPtr.next
		rPtr = rPtr.next
	}

	if lPtr != nil {
		return false, nil
	}

	if rPtr != nil {
		return false, nil
	}

	return true, nil
}

// Find an item in the list.  Return error if not found.
// Implement cmp function object to use this method.
func (list *List) Find(data interface{}, cmp func(interface{}, interface{}) (int, error)) (uint, error) {
	if list.head == nil {
		return 0, notFound()
	}

	var count uint = 0
	current := list.head
	for current != nil {
		b, err := cmp(current.data, data)
		if err != nil {
			return 0, err
		}

		if b == 0 {
			return count, nil
		}

		current = current.next
		count++
	}

	return 0, notFound()
}

// Remove an item from the list.  Return error if not found.
// Implement cmp function object to use this method
func (list *List) Remove(data interface{}, cmp func(interface{}, interface{}) (int, error)) error {
	var previous *node = nil
	var current *node = list.head

	for current != nil {
		b, err := cmp(current.data, data)
		if err != nil {
			return nil
		}

		if b == 0 {
			if current == list.head {
				if list.head == list.tail {
					list.head = nil
					list.tail = nil
				} else {
					list.head = current.next
					list.head.prev = nil
					current.next = nil
					current = nil
				}
			} else if current == list.tail {
				list.tail = current.prev
				list.tail.next = nil
				current.prev = nil
				current = nil
			} else {
				previous.next = current.next
				current.next.prev = previous
				current.prev = nil
				current.next = nil
				current = nil
			}
			return nil
		}

		previous = current
		current = current.next
	}

	return notFound()
}

// Sort the list.  Shallow copy implemented.
// Implement cmp function object to use this method.
func (list *List) Sort(cmp func(interface{}, interface{}) (int, error)) (*List, error) {
	newList := list.Clone()

	// Return the list when 1) the list is empty
	//                      2) the list has only one element
	if newList.head == newList.tail {
		return newList, nil
	}

	ptr := newList.head
	for ptr != nil {
		previous := newList.head
		current := previous.next
		for current != nil {
			c, err := cmp(previous.data, current.data)
			if err != nil {
				return newList, err
			}

			if c > 0 {
				// Swap data
				previous.data, current.data = current.data, previous.data
			}

			previous = current
			current = current.next
		}
		ptr = ptr.next
	}

	return newList, nil
}
