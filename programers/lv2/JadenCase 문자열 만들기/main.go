package main

import (
	"strings"
)

func main() {
}
func solution(s string) string {
	ss := strings.Split(s, "")
	answer := ""
	answer += strings.ToUpper(ss[0])
	for i := 1; i < len(ss); i++ {
		if ss[i] == " " {
			answer += " "
		} else if ss[i-1] == " " {
			answer += strings.ToUpper(ss[i])
		} else {
			answer += strings.ToLower(ss[i])
		}
	}
	return answer
}
