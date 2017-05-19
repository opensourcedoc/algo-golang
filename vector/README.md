# Vector

Vector algebra implemented in Go

## Usage

Vector addition:

```
package main

import (
    "github.com/cwchentw/ds-golang/vector"
    "log"
)

func main() {
    v1 := vector.New(1, 2, 3)
    v2 := vector.New(2.0, 3.0, 4.0)
    v, _ := v1.Add(v2)
    
    n := v.GetAt(1)
    if n.(float64) != 5.0 {
        log.Fatal("Wrong value")
    }
}
```

## Intro

Go lacks vector data structures as those in R and Python. This library demostrates one possible implementation.

## Copyright

2017, Michael Chen

## License

MIT