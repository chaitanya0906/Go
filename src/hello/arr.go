package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[2] = 7
	fmt.Println(a)

	b := [5]int{1, 2, 2, 4, 5}
	fmt.Println(b)

	s := []int{1, 2, 2, 4, 5}
	fmt.Println(s)

	s = append(s, 13)
	fmt.Println(s)

}
