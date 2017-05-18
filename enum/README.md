# Enum

Enumeration implementation in Go

## Usage

Check whether the value exists:

```
package main

import (
    "github.com/cwchentw/ds-golang/enum"
    "log"
)

func main() {
    const color = enum.New("Red", "Green", "Blue")

    if !color.Has("Green") {
	    log.Fatal("Green should be there")
    }
}
```

Get the value from the Enum object:

```
package main

import (
    "fmt"
    "github.com/cwchentw/ds-golang/enum"
)

func main() {
    const color = enum.New("Red", "Green", "Blue")
    
    green := color.Get("Green")
    
    if enum.Eq(green, "Red") {
        fmt.Println("It is red")
    } else if enum.Eq(green, "Green") {
        fmt.Println("It is green")
    } else if enum.Eq(green, "Blue") {
        fmt.Println("It is blue")
    }
}
```

## Intro

Go lacks real enumeration type. Currently, most online tutorials use constant as an alternative. However, that apporach misses real type protecton. The following code is wrongly correct:

```
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

## Copyright

2017, Michael Chen

## License

MIT