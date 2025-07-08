package main

import (
	"fmt"
)

func isDivisibleBy3(n int) bool {
	return n%3 == 0
}
func isDivisibleBy5(n int) bool {
	return n%5 == 0
}

func main() {
	mode := "pinpan" // "divisivelPor3"
	for i := 1; i <= 100; i++ {
		switch mode {
		case "divisivelPor3":
			if isDivisibleBy3(i) {
				fmt.Printf("%d é divisível por 3\n", i)
			} else {
				fmt.Printf("%d não é divisível por 3\n", i)
			}
		case "pinpan":
			if isDivisibleBy3(i) && isDivisibleBy5(i) {
				fmt.Println("Pinpan")
			} else if isDivisibleBy3(i) {
				fmt.Println("Pin")
			} else if isDivisibleBy5(i) {
				fmt.Println("Pan")
			} else {
				fmt.Println(i)
			}
		default:
			fmt.Println("Modo inválido")
			return
		}
	}
}
