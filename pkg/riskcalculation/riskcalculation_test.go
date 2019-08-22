package riskcalculation_test

import (
	"testing"

	"github.com/chanchailee/golang-grpc-credit-risk-api/pkg/riskcalculation"
	"github.com/stretchr/testify/assert"
)

func TestCalFICOScore(t *testing.T) {
	input := map[float64]map[string]float64{
		0.00: {
			"payment_history":          0.0,
			"amount_owned":             0.0,
			"length_of_credit_history": 0.0,
			"new_credit":               0.0,
			"type_of_credit":           0.0,
		},
		1.00: {
			"payment_history":          1.0,
			"amount_owned":             1.0,
			"length_of_credit_history": 1.0,
			"new_credit":               1.0,
			"type_of_credit":           1.0,
		},
		0.34: {
			"payment_history":          0.2,
			"amount_owned":             0.3,
			"length_of_credit_history": 0.5,
			"new_credit":               1.0,
			"type_of_credit":           0.0,
		},
		0.66: {
			"payment_history":          0.7,
			"amount_owned":             0.7,
			"length_of_credit_history": 0.1,
			"new_credit":               1.0,
			"type_of_credit":           0.9,
		},
		0.31: {
			"payment_history":          0.1,
			"amount_owned":             0.1,
			"length_of_credit_history": 0.5,
			"new_credit":               1.0,
			"type_of_credit":           0.7,
		},
	}

	for expected, value := range input {
		actual := riskcalculation.CalFICOScore(value)
		assert.Equal(t, expected, actual)
	}
}
