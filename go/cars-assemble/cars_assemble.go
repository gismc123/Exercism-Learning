package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var percent float64 = successRate / 100
	var pr float64 = float64(productionRate)
	var cwcph float64 = pr * percent
	var answer = cwcph
	return answer
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var a float64 = float64(productionRate)
	var b float64 = successRate / 100
	var c float64 = a / 60
	var d float64 = b * c
	var e int = int(d)
	return e
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	var a int = carsCount
	var numOfTens int = a / 10
	var remainingCars int = a % 10
	var costOfGrouped int = numOfTens * 95000
	var costOfSingles int = remainingCars * 10000
	var answer int = costOfGrouped + costOfSingles
	var b uint = uint(answer)
	// return ((carsCount / 3) * 95000) + (10000 * (carsCount % 3))
	return b
}
