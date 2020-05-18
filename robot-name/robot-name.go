package robotname

import (
	"math/rand"
)

type Robot struct {
	name string
}

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits = "0123456789"

// Name - return random mane for robot (stupid implementation).
// Refactor: apply bites (faster - from 2.621s to 1.535s)
// Bonus test fail in 1883 (not 464)
func (r *Robot) Name() string {
	if r.name != "" {
		return r.name
	}
	b := make([]byte, 5)
	// rand.Int63() produces a random number with 63 random bits.
	for i := 0; i <= 1; i++ {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	for i := 2; i <= 4; i++ {
		b[i] = digits[rand.Int63()%int64(len(digits))]
	}
	r.name = string(b)
	return r.name
}

func (r *Robot) Reset() {
	r.name = ""
}
