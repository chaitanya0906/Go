package main

import "fmt"

func main() {

	vertices := make(map[string]int)

	vertices["triangle"] = 1
	vertices["apple"] = 1
	vertices["Maa"] = 1

	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])
	fmt.Println(vertices["traingl"])

	delete(vertices, "apple")
	fmt.Println(vertices)

}
