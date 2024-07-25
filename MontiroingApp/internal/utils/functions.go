package utils

import "strconv"

func StringTofloat(s string) float32 {

	if len(s) < 4 {
		return 0.0
	}

	usages, _ := strconv.ParseFloat(s[0:4], 32)
	usages = usages / 100
	return float32(usages)
}
