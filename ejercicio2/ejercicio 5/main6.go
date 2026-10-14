package main6

import "fmt"

type CuentaBancaria struct {
	saldo float64
}

func (c *CuentaBancaria) Depositar(monto float64) {
	c.saldo += monto
}

func (c *CuentaBancaria) Retirar(monto float64) {
	if monto <= c.saldo {
		c.saldo -= monto
	} else {
		fmt.Println("Saldo insuficiente")
	}
}

func main() {
	cuenta := CuentaBancaria{saldo: 2000}

	cuenta.Depositar(1000)
	cuenta.Retirar(400)

	fmt.Println("Saldo final:", cuenta.saldo)
}
