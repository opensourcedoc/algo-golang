# Vector

Vector algebra implemented in Go

## Intro

Go lacks vector data structures as those in R and Python. This library demostrates one possible implementation. We are working to shift genercis-based vectors to vectors to specific types because most vectors works on floating-points, occational big decimal. Generics-based vectors seems a bit over-engineering.

## Usage

Vector addition:

```golang
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

```golang
package main

import (
	"github.com/cwchentw/algo-golang/vector/float64"
	"log"
)

func main() {
	v1 := vector.New(1, 2, 3)
	v2 := vector.New(2.0, 3.0, 4.0)
	v := vector.Apply(v1, v2, func(a float64, b float64) float64 { return a + b })

	if v.GetAt(1) != 5.0 {
		log.Fatal("Wrong value")
	}
}
```

## Copyright

2017, Michelle Chen

## License

MIT