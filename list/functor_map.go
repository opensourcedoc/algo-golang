package list

import ()

// Iterate through the list and transform each datum into new value.
// Implement m function object to use this method.
func (list *List) Map(m func(interface{}) (interface{}, error)) (*List, error) {
	newList := New()

	current := list.head
	for current != nil {
		num, err := m(current.data)
		if err != nil {
			return newList, err
		}

		newList.Push(num)

		current = current.next
	}

	return newList, nil
}
