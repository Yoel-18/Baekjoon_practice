package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	aa := float64(a)
	bb := float64(b)
	answer1 := -aa - math.Sqrt(aa*aa-bb)
	answer2 := -aa + math.Sqrt(aa*aa-bb)
	if answer1 == answer2 {
		fmt.Print(answer1)
	} else {
		fmt.Print(answer1, answer2)
	}
	// k := a*a - b
	// d := sort.Search(k, func(i int) bool {
	// 	return i*i >= k
	// })
	// if d == 0 {
	// 	fmt.Println(-a)
	// } else {
	// 	fmt.Println(-a-d, -a+d)
	// }
}
