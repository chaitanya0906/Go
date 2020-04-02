package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	res := sum(2.2, 3)
	fmt.Println(res)

	res2, err := sqrt(-16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res2)
	}

}

func sum(x float64, y int) int {
	return int(x) + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}
