package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println("Hello")
	var frequency int
	for scanner.Scan() {
		sval := scanner.Text()
		fmt.Printf("Adjusting frequency by %s...", sval)
		frequency = Adjust(frequency, sval)
		fmt.Printf("Frequency is %d.\n", frequency)
	}
  fmt.Println("Final Frequency is: %d.\n", frequency)
}

// Adjust takes an int value, and a string in the format "[+-]\d+"
//     and returns the value + / - the string as int.
func Adjust(val int, str string) int {
	adj, _ := strconv.Atoi(str[1:])
	if sval[:1] == "+" {
		return val + adj
	}
	return val - adj
}
