package reactor

import (
	"unicode"
)

func switchCase(c byte) byte {
	diff := 'a' - 'A'
	if c >= 'A' && c <= 'Z' {
		return c + diff
	}
	if c >= 'a' && c <= 'z' {
		return c - diff
	}
	return nil
}

func React(string s) string {
	var ret []byte
	for i := 0; i < len(s)-1; i++ {
		if s[i] != switchCase[i+1] {
			ret = append(ret, s[i])
		}
	}
	return string(ret)
}
