package main

import (
	"strconv"
	"strings"
)

func main() {

}

var (
	check  [10000000]bool
	sel    [1000]bool
	answer int
)

func solution(numbers string) int {
	dfs(0, numbers, "")
	return answer
}
func dfs(count int, s, res string) {
	str := strings.Split(s, "")
	cur, _ := strconv.Atoi(res)
	if res != "" && check[cur] == false {
		num := cur
		check[num] = true
		if isPrime(num) {
			answer++
		}
	}
	for i := 0; i < len(s); i++ {
		if sel[i] == true {
			continue
		}
		sel[i] = true
		dfs(count+1, s, res+str[i])
		sel[i] = false
	}
}
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
