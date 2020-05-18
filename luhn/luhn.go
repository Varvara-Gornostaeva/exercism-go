package luhn

import s "strings"

func Valid(str string) bool {
	if len(str) < 1 || str == "0" {
		return false
	}
	clean := s.TrimSpace(str)
	var sum = 0
	var nDigits = len(clean)
	var parity = nDigits % 2

	for i := 0; i < nDigits; i++ {
		var digit = int(clean[i] - 48)
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0

}