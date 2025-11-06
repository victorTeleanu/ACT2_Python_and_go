// converter/converter_test.go
package converter_test

import (
	"proyecto_go/converter"
	"testing"
)

func almostEqual(a, b float64) bool {
	const eps = 1e-9
	if a > b {
		return a-b < eps
	}
	return b-a < eps
}

// 1️⃣ Test correcto: EUR -> USD
func TestConvertEURtoUSD(t *testing.T) {
	got, _ := converter.Convert(100, "EUR", "USD")
	want := 110.0
	if !almostEqual(got, want) {
		t.Fatalf("100 EUR -> USD = %v; esperaba %v", got, want)
	}
}

// 2️⃣ Test correcto: USD -> EUR
func TestConvertUSDtoEUR(t *testing.T) {
	got, _ := converter.Convert(110, "USD", "EUR")
	want := 100.0
	if !almostEqual(got, want) {
		t.Fatalf("110 USD -> EUR = %v; esperaba %v", got, want)
	}
}

// 3️⃣ Test que falla intencionalmente
func TestFailIntentional(t *testing.T) {
	got, _ := converter.Convert(50, "EUR", "GBP")
	want := 40.0 // valor incorrecto a propósito
	if !almostEqual(got, want) {
		t.Fatalf("50 EUR -> GBP = %v; se esperaba %v", got, want)
	}
}
