package checksum

import (
	"testing"
)

type answer struct {
	twos   bool
	threes bool
}

var (
	testdata = []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	expected = []answer{
		{false, false},
		{true, true},
		{true, false},
		{false, true},
		{true, false},
		{true, false},
		{false, true},
	}
)

func TestChecksum(t *testing.T) {
	for i := range testdata {
        a,b :=Checksum(testdata[i])
		reply := answer{a,b}
		if reply != expected[i] {
			t.Fatalf("Got %v, expected %v", reply, expected[i])
		}

	}

}
