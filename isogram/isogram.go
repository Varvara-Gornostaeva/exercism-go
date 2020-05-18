package main

import (
	s "strings"
)

func IsIsogram(str string) bool {
	// Check is name
	if s.Title(str) == str && len(s.Split(str, " ")) > 1 {
		return true
	}
	// clear string
	formatStr := s.ToLower(s.TrimSpace(str))
	clear := s.Replace(formatStr, "-", "", -1)

	for i := 0; i < len(clear)-1; i++ {
		if s.Count(clear, string(clear[i])) > 1 {
			return false
		}
	}
	return true
}
