package main

import (
	"fmt"
	"time"
)

func miFuncion() {
	fmt.Println("¡Hello word")
}

func main() {
	go miFuncion()
	time.Sleep(1 * time.Second)
}
