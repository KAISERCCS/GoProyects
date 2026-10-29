package main

import (
	"fmt"
	"sync"
)

func main() {
	entrada := make(chan int)
	resultados := make(chan int)

	// Generador de números
	go func() {
		for i := 1; i <= 10; i++ {
			entrada <- i
		}
		close(entrada)
	}()

	// Fan-out: varios trabajadores
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			for numero := range entrada {
				cuadrado := numero * numero
				fmt.Println("Trabajador", id, "procesó", numero)
				resultados <- cuadrado
			}
		}(i)
	}

	// Fan-in: reunir los resultados
	go func() {
		wg.Wait()
		close(resultados)
	}()

	// Mostrar resultados finales
	fmt.Println("Cuadrados obtenidos:")

	for resultado := range resultados {
		fmt.Println(resultado)
	}
}
