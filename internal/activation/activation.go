package activation

func Activation(output []float64) string {
	var max float64 = -99999
	pos := -1

	for i, val := range output {
		if val > max {
			max = val
			pos = i
		}
	}

	switch pos {
	case 0:
		return "Up"
	case 1:
		return "Down"
	}

	return ""
}
