package main

import "fmt"

func main() {
	hecho := make(chan bool)

	go func() {
		suma := 0

		for i := 1; i <= 10; i++ {
			suma += i
		}

		fmt.Println("La suma es:", suma)
		hecho <- true
	}()

	<-hecho
	fmt.Println("Programa terminado")
}
