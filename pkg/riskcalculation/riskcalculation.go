package riskcalculation

import "math"

var riskFactorWeight = map[string]float64{
	"payment_history":          0.35,
	"amount_owned":             0.30,
	"length_of_credit_history": 0.15,
	"new_credit":               0.10,
	"type_of_credit":           0.10,
}

// CalFICOScore function is used to calculate FICO Score by using weight mean formula.
// It receives 	map[string]float64 	as a parameter and return a score as a float64 with 2 decimals.
func CalFICOScore(input map[string]float64) float64 {
	var ficoScore float64

	for key := range input {
		score := input[key] * riskFactorWeight[key]
		ficoScore += score
	}

	return math.Round(ficoScore*100) / 100
}
