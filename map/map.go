package main

import "fmt"

func main() {
	x:= make (map[string]int)
	x["a"]= 20
	x["b"]= 30
	fmt.Println(x["a"])
}