# list

An implementation of doubly-linked list in Go

## Usage

A doubly-linked list implementation in Go

This library implements doubly-linked list in Go. Since Go doesn't support generics, this library utilize empty interface to emulate this functionality, delegating element comparison tasks to users.

```
package main

import (
    "github.com/cwchentw/ds-golang/list"
    "log"
)

func main() {
    l := list.New()

    for i := 1; i <= 10; i++ {
        l.Push(i)
    }

    data, _ := l.At(3)
    if data != 4 {
        log.Fatalln("Data error")
    }
}
```

This library implements higher order functions as well. Same, the library delegates related interface implementations to users.

```
package main

import (
    "errors"
    "github.com/cwchentw/ds-golang/list"
    "log"
)

func main() {
    l := list.New()

    for i := 1; i <= 10; i++ {
        l.Push(i)
    }

    evens, _ := l.Select(even)
    for e := range evens.Iter() {
        n, _ := e.(int)
	    if n % 2 != 0 {
            log.Fatalln("Error Select result")
        }
    }
}

func even(a interface{}) (bool, error) {
    n, ok := a.(int)
    if ok != true {
        return false, errors.New("Failed type assertion on a")
    }

    return n % 2 == 0, nil
}
```

## Copyright

2017, Michael Chen

## License

MIT