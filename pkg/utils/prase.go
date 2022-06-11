package utils

import "strconv"

func ParseStr2Uint(str string) (uint, bool) {
	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, false
	}
	return uint(id), true
}
