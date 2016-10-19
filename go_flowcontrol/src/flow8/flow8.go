package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := math.Sqrt(x)
    return z
}

func main() {
    fmt.Println(Sqrt(144))
}