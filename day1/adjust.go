package main

import (
	"strconv"
)

// Adjust takes an int value, and a string in the format "[+-]\d+"
//     and returns the value + / - the string as int.
func Adjust(val int, str string) int {
	adj, _ := strconv.Atoi(str[1:])
	if str[:1] == "+" {
		return val + adj
	}
	return val - adj
}
