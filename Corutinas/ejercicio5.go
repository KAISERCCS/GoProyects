package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			fmt.Println("Goroutine", id, "terminada")
		}(i)
	}

	wg.Wait()
	fmt.Println("Todas las goroutines terminaron")
}
