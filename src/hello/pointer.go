package main

import "fmt"

func main() {
	i := 7
	fmt.Println("Mem Addr", &i)

	inc1(i)

	fmt.Println(i)
	inc2(&i)
	fmt.Println(i)

}

func inc1(i int) {
	i++
}

func inc2(i *int) {
	(*i)++
}
