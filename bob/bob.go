// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	s "strings"
	"unicode"
)

const (
	yell         = "Whoa, chill out!"
	question     = "Sure."
	yellquestion = "Calm down, I know what I'm doing!"
	empty        = "Fine. Be that way!"
	defaultBob   = "Whatever."
)

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	cleanString := s.TrimSpace(remark)
	if len(cleanString) == 0 {
		return empty
	}
	IsLetterText, _ := regexp.MatchString("[a-zA-Z]+", cleanString)
	isYell := IsLetterText && cleanString == s.ToUpper(cleanString)
	isQuestion := s.HasSuffix(cleanString, "?")
	if isYell && isQuestion {
		return yellquestion
	}
	if isQuestion {
		return question
	}
	if isYell {
		return yell
	}

	return defaultBob
}
