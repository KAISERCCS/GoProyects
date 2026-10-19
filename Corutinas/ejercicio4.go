package main

import "fmt"

func sumar(a int, b int, ch chan int) {
	resultado := a + b
	ch <- resultado
}

func main() {
	ch := make(chan int)

	go sumar(5, 3, ch)

	fmt.Println("Resultado:", <-ch)
}
