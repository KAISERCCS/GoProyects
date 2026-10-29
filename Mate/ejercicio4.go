package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	const goroutines = 10
	const repeticiones = 100000

	// Ejemplo sin protección
	var contador int
	var wg sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < repeticiones; j++ {
				contador++
			}
		}()
	}

	wg.Wait()
	fmt.Println("Sin protección:", contador)
	fmt.Println("Resultado esperado:", goroutines*repeticiones)

	// Solución con Mutex
	contador = 0
	var mutex sync.Mutex

	for i := 0; i < goroutines; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < repeticiones; j++ {
				mutex.Lock()
				contador++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Con Mutex:", contador)

	// Solución con atomic
	var contadorAtomic int64
	var wgAtomic sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wgAtomic.Add(1)

		go func() {
			defer wgAtomic.Done()

			for j := 0; j < repeticiones; j++ {
				atomic.AddInt64(&contadorAtomic, 1)
			}
		}()
	}

	wgAtomic.Wait()
	fmt.Println("Con Atomic:", contadorAtomic)
}
