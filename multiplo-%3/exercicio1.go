package main

import "fmt"


func main() {
    for i := 0; i <= 100; i++ {
        if i%3 == 0 {
            fmt.Println("PIN")
        } else if i%5 == 0 {
            fmt.Println("PAN")
    }
}
}