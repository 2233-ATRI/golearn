package gotest

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, nil
	}
	return a / b, nil
}
