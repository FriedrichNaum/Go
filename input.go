package main

import (
    "fmt"
)

// This is a simple Go program
func main() {
    var x int = 42
    y := 3.14
    fmt.Println("Hello, World!", x, y)
    
    if x > 10 {
        fmt.Println("x is greater than 10")
    } else {
        fmt.Println("x is 10 or less")
    }
}