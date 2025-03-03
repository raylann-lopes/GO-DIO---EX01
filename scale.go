package main

import "fmt"

const ebuKelvin = 373

func main() {
	celsius := ebuKelvin - 273
	fmt.Println("O ponto de ebulição da agua em Kelvin é de" , ebuKelvin, "e o ponto de ebulição da agua em Celsius é de", celsius,"!")
}