package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n, sum := 0, 0
	var str string
	fmt.Scan(&n)
	fmt.Scan(&str)
	num := strings.Split(str, "")
	for i := 0; i < n; i++ {
		number, _ := strconv.Atoi(num[i])
		sum += number
	}
	fmt.Print(sum)
}
