package math

func Abs(x float64) float64 {
	if x >= 0 {
		return x
	}

	return x * -1.0
}
