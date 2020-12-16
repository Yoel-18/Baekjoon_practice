package main

import (
	"sort"
	"strconv"
	"strings"
)

func main() {
	solution("1 2 3 4")
}
func solution(s string) string {
	str := strings.Split(s, " ")
	num := make([]int, 0)
	for i := 0; i < len(str); i++ {
		number, _ := strconv.Atoi(str[i])
		num = append(num, number)
	}
	sort.Ints(num)
	min := strconv.Itoa(num[0])
	max := strconv.Itoa(num[len(num)-1])
	answer := min + " " + max
	return answer
}

// func solution(s string) string {
//     strs := strings.Fields(s)

//     ints := make([]int, len(strs))
//     for i, str := range strs {
//         ints[i], _  = strconv.Atoi(str)
//     }
//     sort.Ints(ints)

//     return fmt.Sprintf("%d %d", ints[0], ints[len(ints)-1])
// }
