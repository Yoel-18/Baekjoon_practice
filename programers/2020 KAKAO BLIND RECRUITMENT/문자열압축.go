package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Print(solution(str))
}
func solution(str string) int {
	answer := len(str) // 초기값 문자열 길이
	//	1,2,3... -> 한번에 끊을 문자열, 최대 절반까지
	for i := 1; i < len(str)/2; i++ {
		var temp string
		//	처음부터 i씩 귾어서 비교
		for j := 0; j < len(str); j += i {
			cur := ""
			if i+j >= len(str) {
				cur = str[j:len(str)]
			} else {
				cur = str[j : j+i]
			}
			count := 1
			var sb string

			for k := i + j; k < len(str); k += i {
				cur2 := ""

				if k+i >= len(str) {
					cur2 = str[k:len(str)]
				} else {
					cur2 = str[k : k+i]
				}

				if cur == cur2 {
					count++
					j = k
				} else {
					break
				}
			}
			if count == 1 {
				sb = cur
			} else {
				sb = strconv.Itoa(count) + cur
			}
			temp = sb
		}
		if answer > len(temp) {
			answer = len(temp)
		}
	}
	return answer
}
