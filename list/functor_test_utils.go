package list

func newEvenList() *List {
	list := New()

	list.Unshift(100)
	list.Unshift(200)
	list.Unshift(300)
	list.Unshift(400)

	return list
}

func newOddList() *List {
	list := New()

	list.Unshift(111)
	list.Unshift(333)
	list.Unshift(555)
	list.Unshift(777)

	return list
}

func new1To10List() *List {
	list := New()

	for i := 1; i <= 10; i++ {
		list.Push(i)
	}

	return list
}
