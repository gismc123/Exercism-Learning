// Package weather calculates the weather.
package weather

// CurrentCondition describes the current weather.
var CurrentCondition string

// CurrentLocation tells the city that the weather is affecting.
var CurrentLocation string

// Forecast returns the cities current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
