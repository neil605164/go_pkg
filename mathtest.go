package mathtest

// Average 計算平均
func Average(xs []float64) float64 {
	total := float64(0)

	for k := range xs {
		total += xs[k]
	}

	return total / float64(len(xs))
}

// Add 計算相加
func Add(xs []float64) float64 {
	total := float64(0)

	for k := range xs {
		total += xs[k]
	}

	return total
}

// Negative 計算相減
func Negative(xs []float64) float64 {
	total := float64(0)

	for k := range xs {
		total -= xs[k]
	}

	return total
}

// PrintSomething 字串組合
func PrintSomething(str ...string) string {
	strSelf := ""
	for k := range str {
		strSelf += str[k]
	}

	return strSelf
}
