package main

import (
	"fmt"
	"log"
	"proyecto_go/converter"
)

func main() {
	amount := 100.0
	from := "EUR"
	to := "USD"

	converted, err := converter.Convert(amount, from, to)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%.2f %s = %.2f %s\n", amount, from, converted, to)
}
