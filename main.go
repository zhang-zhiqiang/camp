package main

import "fmt"

func main() {
    a := "test"
    b := &a
    c := *&a
    fmt.Printf("%s\n", a)
    fmt.Printf("%v\n", b)
    fmt.Printf("%v\n", c)
    a = "99999"
    fmt.Printf("%s\n", a)
    fmt.Printf("%v\n", b)
    fmt.Printf("%v\n", c)
}
