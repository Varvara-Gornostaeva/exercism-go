package raindrops

import "fmt"

func GetFactors(nr int64) {
	fmt.Printf("\nFactors of %d: ", nr)
	fs := make([]int64, 1)
	fs[0] = 1
	apf := func(p int64, e int) {
		n := len(fs)
		for i, pp := 0, p; i < e; i, pp = i+1, pp*p {
			for j := 0; j < n; j++ {
				fs = append(fs, fs[j]*pp)
			}
		}
	}
	e := 0
	for ; nr & 1 == 0; e++ {
		nr >>= 1
	}
	apf(2, e)
	for d := int64(3); nr > 1; d += 2 {
		if d*d > nr {
			d = nr
		}
		for e = 0; nr%d == 0; e++ {
			nr /= d
		}
		if e > 0 {
			apf(d, e)
		}
	}
	return fs
}

func Convert(number int) string {

}