package main

type retangulo struct {
	comprimento, altura int
}

func (r retangulo) area() int {
	return r.comprimento * r.altura
}

func (r retangulo) perimetro() int {
	return 2 * (r.comprimento + r.altura)
}

func main() {
	r := retangulo{comprimento: 10, altura: 5}
	println(r.area())
	println(r.perimetro())
}
