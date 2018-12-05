package main

import (
	"aoc18/day2/checksum"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hello\n")

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var twee, drie int
	var ids []string
	for scanner.Scan() {
		text := scanner.Text()
		ids = append(ids, text)
		a, b := checksum.Checksum(text)
		if a {
			twee++
		}
		if b {
			drie++
		}
	}
	fmt.Printf("Final checksum: %d\n", twee*drie)
	fmt.Printf("Common letters: %s\n", checksum.Overlap(ids))
}
