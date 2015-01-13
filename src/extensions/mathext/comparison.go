package mathext

func Min(lessThan func(x, y interface{}) bool, data ...interface{}) interface{} {
	l := len(data)
	if l == 0 {
		return nil
	}
	min := data[0]
	for i := 1; i < l; i++ {
		if lessThan(data[i], min) {
			min = data[i]
		}
	}
	return min
}

func Max(lessThan func(x, y interface{}) bool, data ...interface{}) interface{} {
	l := len(data)
	if l == 0 {
		return nil
	}
	max := data[0]
	for i := 1; i < l; i++ {
		if !lessThan(data[i], max) {
			max = data[i]
		}
	}
	return max
}
