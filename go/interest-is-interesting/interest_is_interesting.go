package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch{
	case balance < 0:
		return float32(3.213)
	case balance < 1000:
		return float32(0.5)
	case balance < 5000:
		return float32(1.621)
	case balance >= 5000:
		return float32(2.475)
	default:
		return float32(0)
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	rate := float64(InterestRate(balance))/100
	interest := balance * rate
	return float64(interest)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return float64(balance + Interest(balance))
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	numberOfYears := 0
	startingBalance := float64(balance)
	for startingBalance < targetBalance {
		startingBalance = AnnualBalanceUpdate(startingBalance)
		numberOfYears++
	}
	return numberOfYears
}
