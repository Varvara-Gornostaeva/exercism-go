package distance

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	distance := 0
	if a == b {
		return 0, nil
	}
	if len(a) != len(b) {
		return -1, errors.New(" must be the same length")
	}
	for i, j := 0, 0; i < len(a); i, j = i+1, j+1 {
		if a[i] != b[i] {
			distance += 1
		}
	}
	return distance, nil

}