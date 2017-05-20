package list

import ()

// Iterate through the list and transform each datum into new value.
// Implement m function object to use this method.
func (list *List) Map(m func(interface{}) (interface{}, error)) (*List, error) {
	newList := New()

	_len := list.Len()
	cnum := make(chan interface{}, _len)
	cerr := make(chan error, _len)
	current := list.head
	for current != nil {
		go func() {
			d, e := m(current.data)
			cnum <- d
			cerr <- e
		}()

		err := <-cerr
		if err != nil {
			return newList, err
		}
		newList.Push(<-cnum)

		current = current.next
	}

	return newList, nil
}
