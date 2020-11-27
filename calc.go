package calc

func Add(addend1, addend2 int) int {
	return addend1 + addend2
}

func Subtract(minuend, subtrahend int) int {
	return minuend - subtrahend
}

func Multiply(multiplicant, multiplier int) int {
	return multiplicant * multiplier
}

func Divide(dividend, divisor int) float64 {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	return float64(dividend) / float64(divisor)
}
