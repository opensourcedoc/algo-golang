# Vector

Vector algebra implemented in Go

## Intro

Go lacks vector data structures as those in R and Python. This library demostrates one possible implementation. We are working to shift genercis-based vectors to vectors specific type because most vectors works on floating-point, occational big decimal.

## Usage

Vector addition:

```
package main

import (
	"github.com/cwchentw/algo-golang/vector/float64"
	"log"
)

func main() {
	v1 := vector.New(1, 2, 3)
	v2 := vector.New(2.0, 3.0, 4.0)
	v := vector.Add(v1, v2)

	n := v.GetAt(1)
	if n != 5.0 {
		log.Fatal("Wrong value")
	}
}
```

Vector algebra by delegation:

```
package main

import (
    "errors"
    "github.com/cwchentw/algo-golang/vector/float64"
    "log"
)

func main() {
    v1 := vector.New(1, 2, 3)
    v2 := vector.New(2.0, 3.0, 4.0)
    v, _ := v1.Apply(v2, add)
    
    n := v.GetAt(1)
    if n.(float64) != 5.0 {
        log.Fatal("Wrong value")
    }
}

func add(a interface{}, b interface{}) (interface{}, error) {
    na, ok := a.(int)
    if !ok {
        return 0.0, errors.New("Wrong value a")
    }
    
    nb, ok := b.(float64)
    if !ok {
        return 0.0, errors.New("Wrong value b")
    }
    
    return float64(a) + b, nil
}
```

## Copyright

2017, Michael Chen

## License

MIT