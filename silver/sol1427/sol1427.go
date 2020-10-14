package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	numbers := strings.Split(str, "")
	num := make([]int, 0)
	for _, number := range numbers {
		k, _ := strconv.Atoi(number)
		num = append(num, k)
	}
	sort.Ints(num)
	for i := len(num) - 1; i >= 0; i-- {
		fmt.Print(num[i])
	}
}
