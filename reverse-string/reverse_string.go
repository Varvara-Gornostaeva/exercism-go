package reverse

func String(orig string) string {
	// Using string - tests time 2.883s
	// Using rune - test time 1.744s
	r := []rune(orig)
	j := len(r) - 1
	for i := 0; i < j; i++ {
		r[i], r[j] = r[j], r[i]
		j--
	}
	return string(r)
}
