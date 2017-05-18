package list

// Check whether any item in the list fulfills the predicate.
// Implement grep function object to use this method.
func (list *List) Any(grep func(interface{}) (bool, error)) (bool, error) {
	current := list.head
	for current != nil {
		b, err := grep(current.data)
		if err != nil {
			return false, err
		}

		if b == true {
			return true, nil
		}

		current = current.next
	}

	return false, nil
}

// Check whether all items in the list fulfill the predicate.
// Implement grep function object to use this method.
func (list *List) All(grep func(interface{}) (bool, error)) (bool, error) {
	current := list.head

	for current != nil {
		b, err := grep(current.data)
		if err != nil {
			return false, err
		}

		if b != true {
			return false, nil
		}

		current = current.next
	}

	return true, nil
}

// Select items from the list when item fulfill the predicate.
// Implement grep function object to use this method.
func (list *List) Select(grep func(interface{}) (bool, error)) (*List, error) {
	newList := New()

	current := list.head
	for current != nil {
		b, err := grep(current.data)
		if err != nil {
			return newList, err
		}

		if b == true {
			newList.Push(current.data)
		}

		current = current.next
	}

	return newList, nil
}
