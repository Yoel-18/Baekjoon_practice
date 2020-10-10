package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	m := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	day := [7]string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}
	for i := 0; i < x-1; i++ {
		y += m[i]
	}
	fmt.Println(day[y%7])
}
