package main

import "fmt"

func main() {
    defer fmt.Println("world") // menunda eksekusi dgn defer

    fmt.Println("hello")
}