package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Kapan hari Sabtu?")

    today := time.Now().Weekday()

    switch time.Saturday {
        case today + 0:
            fmt.Println("Sekarang.")    
        case today + 1:
            fmt.Println("Satu hari lagi.")
        case today + 2:
            fmt.Println("Dua hari lagi.")
        default:
            fmt.Println("Masih Jauh.")            
    }
}