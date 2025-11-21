package main

import "fmt"

func readInput() (float64, string, string) {
	var amount float64
	var from string
	var to string
	fmt.Print("Введите сумму: ")
	fmt.Scanln(&amount)
	fmt.Print("Введите исходную валюту: ")
	fmt.Scanln(&from)
	fmt.Print("Введите целевую валюту: ")
	fmt.Scanln(&to)
	return amount, from, to
}

func Convert(amount float64, from string, to string) float64 {

	return 0
}

func main() {
	amount, from, to := readInput()
	result := Convert(amount, from, to)
	fmt.Println(result)
}
