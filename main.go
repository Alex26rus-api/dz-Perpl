package main

import "fmt"

func main() {

	const USD_TO_EUR = 0.92
	const USD_TO_RUB = 95.50

	eurToRub := (1 / USD_TO_EUR) * USD_TO_RUB

	fmt.Printf("1 EUR = %.2f RUB\n", eurToRub)

}
