package main

import "fmt"

func main() {
	x := 5
	if x > 6 {
		fmt.Printf("G6")
	} else if x < 2 {
		fmt.Printf("L2")
	} else {
		fmt.Printf("LE6GE2")
	}
}
