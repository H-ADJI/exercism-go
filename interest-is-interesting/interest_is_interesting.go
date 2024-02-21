package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {

	case balance < 0:
		return 3.213
	case balance < 1_000:
		return 0.5
	case balance < 5_000:
		return 1.621
	default:
		return 2.475
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	n_years := 0
	if balance <= 0 || balance >= targetBalance {
		return 0
	}
	for growth := balance; targetBalance-growth > 0; growth = AnnualBalanceUpdate(growth) {
		n_years += 1

	}
	return n_years
}
