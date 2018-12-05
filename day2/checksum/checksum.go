package checksum

import ()

func Checksum(s string) (twos, threes bool) {
	listed := make(map[rune]int)
	for _, rune := range s {
		listed[rune] += 1
	}
	for _, v := range listed {
		if v == 2 {
			twos = true
		}
		if v == 3 {
			threes = true
		}
	}
	return
}

const allowed int = 1

func Overlap(s []string) string {
	for _, v := range s {
		for _, w := range s {
			if v == w {
				continue
			}
			var diff int
			for i := 0; i < len(v); i++ {
				if v[i] == w[i] {
					continue
				}
				diff++
				if diff > allowed {
					break
				}
			}
			if diff == allowed {
				return getCommon(v, w)
			}
		}
	}
	return ""
}

func getCommon(alpha, beta string) string {
	if len(alpha) != len(beta) {
		return ""
	}
	var ret []byte
	for i := 0; i < len(alpha); i++ {
		if alpha[i] == beta[i] {
			ret = append(ret, alpha[i])
		}
	}
	return string(ret)
}
