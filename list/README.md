# List

Doubly-linked list in Go

## Intro

A doubly-linked list implementation in Go

This library implements doubly-linked list in Go. Since Go doesn't support generics, this library utilize empty interface to emulate this functionality, delegating element comparison tasks to users.

This library implements common higher order functions as well. Same, the library delegates related function objects to users.


## Usage

Find specific elements:

```
package main

import (
    "github.com/cwchentw/algo-golang/list"
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

Select only even elements:

```
package main

import (
    "errors"
    "github.com/cwchentw/algo-golang/list"
    "log"
)

func main() {
    l := list.New()

    for i := 1; i <= 10; i++ {
        l.Push(i)
    }

    evens, _ := l.Select(isEven)
    for e := range evens.Iter() {
        n, _ := e.(int)
	    if n % 2 != 0 {
            log.Fatalln("Error Select result")
        }
    }
}

func isEven(a interface{}) (bool, error) {
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