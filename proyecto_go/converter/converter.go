package converter

import "fmt"

// Tasas de cambio respecto al euro
var rates = map[string]float64{
	"EUR": 1.0,
	"USD": 1.10,
	"GBP": 0.85,
}

// Convert convierte una cantidad de `from` a `to`
func Convert(amount float64, from, to string) (float64, error) {
	fromRate, ok := rates[from]
	if !ok {
		return 0, fmt.Errorf("moneda desconocida: %s", from)
	}
	toRate, ok := rates[to]
	if !ok {
		return 0, fmt.Errorf("moneda desconocida: %s", to)
	}
	// Convertimos a euros primero, luego a la moneda destino
	inEUR := amount / fromRate
	return inEUR * toRate, nil
}
