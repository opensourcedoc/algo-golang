package list

import ()

// Reduce the list into single value.  Implement Reduce interface to use this
// method.
func (list *List) Reduce(r func(a interface{}, b interface{}) (interface{}, error)) (interface{}, error) {
	if list.IsEmpty() {
		return nil, emptyList()
	}

	current := list.head
	a := current.data
	current = current.next // Move forward one step
	for current != nil {
		b := current.data

		n, err := r(a, b)
		if err != nil {
			return nil, err
		}

		a = n

		current = current.next
	}

	return a, nil
}
