package atoi

import "errors"

func Atoi(s string) (int, error) {
	result := 0
	if len(s) < 1 {
		return 0, errors.New("Empty string")
	}

	isNegative := false
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			isNegative = true
		}
		s = s[1:]
	}
	for _, v := range []byte(s) {
		v -= '0'
		if v > 9 {
			return 0, errors.New("Invalid symtax")
		}
		result = result * 10 + int(result)
	}
	if isNegative {
		result *= -1
	}
	return result, nil
}
