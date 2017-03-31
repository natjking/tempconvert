package tempconvert

import (
	"math"
)

func SimpleRound(f float64) float64 {
	return math.Floor(f + .5)
}

func Round(f float64, places int) (float64) {
	shift := math.Pow(10, float64(places))
	return SimpleRound(f * shift) / shift;
}

func FahrenheitToCelsius(intemp float64) (float64) {
	return Round((intemp - 32)*5/9, 5)
}

func CelsiusToFahrenheit(intemp float64) (float64) {
	return Round(intemp*9/5+32, 5)
}