package main

import "fmt"

func main() {
	// for
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0

	// while
	for i < 5 {
		fmt.Println(i)
		i++
	}
	arr := [3]string{"a", "b", "c"}

	// auto
	for idx, val := range arr {
		fmt.Println("index:", idx, "value:", val)
	}

	m := make(map[string]string)

	m["a"] = "awfaf"
	m["b"] = "bwfbf"

	// auto
	for idx, val := range m {
		fmt.Println("index:", idx, "value:", val)
	}

}
