package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	parttwo := true

	fmt.Println("Hello")
	var freq int
	found := make(map[int]struct{})

	for {
		file, err := os.Open("input.txt")
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			sval := scanner.Text()
			fmt.Printf("Current frequency %d, change of %s;", freq, sval)
			freq = Adjust(freq, sval)
			if _, ok := found[freq]; ok && parttwo {
				fmt.Printf("Reached %d twice.\n", freq)
				os.Exit(0)
			}
			found[freq] = struct{}{}
			fmt.Printf("resulting frequency %d.\n", freq)
			if !parttwo {
				break
			}
		}
	}
	fmt.Printf("Final Frequency is: %d.\n", freq)
}
