package main

import "fmt"

func main() {
	st := make([]int, 0)
	st = append(st, 1)
	st = append(st, 2)
	fmt.Println(len(st))
	st = append(st, 3)
	fmt.Println(st)
	st = append(st, 4)
	fmt.Println(st)
	fmt.Println(len(st))
	st = st[:2]
	fmt.Println(st)
	fmt.Println(len(st))
}
