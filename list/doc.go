// A doubly-linked list implementation in Go
//
// This library implements doubly-linked list in Go.  Since Go doesn't
// support generics, this library utilize empty interface to mimic this
// functionality and element comparison tasks are delegated to users.
//
// Example:
//
//     package main
//
//     import (
//         "github.com/cwchentw/golist"
//         "log"
//     )
//
//     func main() {
//         list := list.New()
//
//         for i := 1; i <= 10; i++ {
//             list.Push(i)
//         }
//
//         data, _ := list.At(3)
//         if data != 4 {
//             log.Fatalln("Data error")
//         }
//     }
//
// This library implements higher order functions as well.  As Go lacks
// generics, the library delegates related interface implementations
// to users.
//
// Example:
//
//     import (
//         "errors"
//         "github.com/cwchentw/golist"
//         "log"
//     )
//
//     func main() {
//         list := list.New()
//
//         for i := 1; i <= 10; i++ {
//             list.Push(i)
//         }
//
//         evens, _ := list.Select(even)
//         for e := range evens.Iter() {
//             n, _ := e.(int)
//             if n % 2 != 0 {
//                 log.Fatalln("Error Select result")
//             }
//         }
//     }
//
//     func even(a interface{}) (bool, error) {
//         n, ok := a.(int)
//         if ok != true {
//             return false, errors.New("Failed type assertion on a")
//         }
//
//         return n % 2 == 0, nil
//     }
package list
