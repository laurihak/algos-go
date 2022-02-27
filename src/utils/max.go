package utils

func Max(v []int) int {
	m := 0
	for i, e := range v {
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}

func Min(v []int) int {
	m := 0
	for i, e := range v {
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}
