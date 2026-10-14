package main3

import "fmt"

func intercambiar(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a := 5
	b := 10

	fmt.Println("Antes:", a, b)

	intercambiar(&a, &b)

	fmt.Println("Despues:", a, b)
}
