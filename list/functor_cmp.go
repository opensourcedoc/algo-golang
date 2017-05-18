package list

// Check whether two lists are equal.  If both the lengths and data of the
// two list are identical, the two lists are equal.
// Implement eq function object to use this method.
func (list *List) Equal(other *List, eq func(interface{}, interface{}) (bool, error)) (bool, error) {
	lPtr := list.head
	rPtr := other.head

	for lPtr != nil && rPtr != nil {
		b, err := eq(lPtr.data, rPtr.data)
		if err != nil {
			return false, err
		}

		if !b {
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
func (list *List) Find(data interface{}, cmp func(interface{}, interface{}) (bool, error)) (uint, error) {
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

		if b {
			return count, nil
		}

		current = current.next
		count++
	}

	return 0, notFound()
}

// Remove an item from the list.  Return error if not found.
// Implement cmp function object to use this method
func (list *List) Remove(data interface{}, cmp func(interface{}, interface{}) (bool, error)) error {
	var previous *node = nil
	var current *node = list.head

	for current != nil {
		b, err := cmp(current.data, data)
		if err != nil {
			return nil
		}

		if b {
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
// Implement pred function object to use this method.
func (list *List) Sort(pred func(interface{}, interface{}) (bool, error)) (*List, error) {
	newList := New()

	ptr := list.head
	for ptr != nil {
		if newList.IsEmpty() {
			newList.Push(ptr.data)
		} else {
			var p *node = nil
			q := newList.head

			isAdded := false
			for q != nil {
				b, err := pred(ptr.data, q.data)
				if err != nil {
					return newList, err
				}

				if b {
					if q == newList.head {
						newList.Unshift(ptr.data)
						isAdded = true
						break
					} else {
						n := node{data: ptr.data, prev: nil, next: nil}
						n.prev = p
						n.next = q
						q.prev = &n
						p.next = &n
						isAdded = true
						break
					}
				}

				p = q
				q = q.next
			}

			if !isAdded {
				newList.Push(ptr.data)
			}
		}

		ptr = ptr.next
	}

	return newList, nil
}
