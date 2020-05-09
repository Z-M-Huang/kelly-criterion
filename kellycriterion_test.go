package kellycriterion

import (
	"math"
	"testing"
)

func TestSimulate(t *testing.T) {
	winChance := float64(0.6)
	payout := float64(1)
	result := float64(0.2)

	calc := Calculate(winChance, payout)

	if math.Abs(calc-result) > 1e-9 {
		t.Errorf("Expected result is %.02f, got %.02f", result, calc)
	}
}
