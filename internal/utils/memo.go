package utils

func GenerateEasinessFactor(lastEasinessFactor float64, lossFactor float64) float64 {
	if (lastEasinessFactor * float64(lossFactor)) <= 0 {
		return 1
	} else {
		return lastEasinessFactor * float64(lossFactor)
	}
}

func CalculatePeriod(review int, easinessFactor float64) float64 {
	if review == 1 {
		return 1
	} else if review == 2 {
		return 6
	} else {
		return CalculatePeriod(review-1, easinessFactor) * easinessFactor
	}
}
