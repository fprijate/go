package test1

func Add(a, b float64) float64 {
	return a + b
}

func Div(a, b float64) float64 {
	if b != 0 {
		return a * b
	}
	return 0
}
