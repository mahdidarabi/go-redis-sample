package utils

import "strconv"

func ParseStringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}

func ParseStringToBool(str string) bool {
	return str == "true"
}
