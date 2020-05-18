// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	sClean:= strings.Replace(s, "-", " ", -1)
	sSplit:= strings.Split(sClean, " ")
	var abvv []string
	for i:=0; i < len(sSplit); i++ {
		word:= string(sSplit[i][0:1])
		abvv = append(abvv, strings.ToUpper(word))
	}
	return strings.Join(abvv, "")
}