# Enum

Enumeration implementation in Go

## Intro

Go lacks real enumeration type. Currently, most online tutorials use constant as an alternative. However, that apporach misses real type protecton. The following code is wrongly correct:

```golang
package main

import "log"

type Color int

const (
    Red Color = iota + 1
    Green
    Blue
)

func main() {
    if !(1 == Red) {
        log.Fatal("It is not equal")
    }
}
```

In our approach, we use set (as math) to emulate enumeration. When programmers call non-existing values, our program will emit some error to notify users.

## Usage

Check whether the value exists:

```golang
package main

import (
    "github.com/cwchentw/ds-golang/enum"
    "log"
)

func main() {
    var color = enum.New("Red", "Green", "Blue")

    if !color.Has("Green") {
	    log.Fatal("Green should be there")
    }
}
```

Get the value from the Enum object:

```golang
package main

import (
    "fmt"
    "github.com/cwchentw/ds-golang/enum"
)

func main() {
    var color = enum.New("Red", "Green", "Blue")

    green := color.Get("Green")

    if color.Eq(green, "Red") {
        fmt.Println("It is red")
    } else if color.Eq(green, "Green") {
        fmt.Println("It is green")
    } else if color.Eq(green, "Blue") {
        fmt.Println("It is blue")
    }
}
```

## Copyright

2017, Michelle Chen

## License

MIT
