package atoi

import "errors"

func Atoi(s string) (int,error)  {
	result := 0
	if len(s) == 0 {
		return 0, errors.New("Empty string")
	}
	return result,nil
}
