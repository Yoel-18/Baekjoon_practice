package main

import (
	"fmt"
)

func main() {
	fmt.Print(solution("JAN"))
}

// func solution(s string) string {
// 	str := strings.ToLower(s)
// 	buf := make([]byte, 0)
// 	for i := 0; i < len(str); i++ {
// 		if i == 0 && str[i] >= 'a' && str[i] <= 'z' {
// 			buf = append(buf, str[i]-32)
// 			continue
// 		}
// 		if i != 0 && str[i-1] == ' ' && str[i] >= 'a' && str[i] <= 'z' {
// 			buf = append(buf, str[i]-32)
// 			continue
// 		}
// 		buf = append(buf, str[i])
// 	}
// 	return string(buf)
// }

//	a b c d e f g h i j k l m
//  n o p q r s t u t w x y z
func solution(name string) int {
	count := 0
	move := len(name) - 1
	for i := 0; i < len(name); i++ {
		if name[i] != 'A' {
			up := int(name[i] - 'A')
			down := (int)('Z' - name[i] + 1)
			count += minInt(up, down)
		}
	}
	for i := 0; i < len(name); i++ {
		if name[i] != 'A' {
			next := i + 1
			for next < len(name) && name[next] == 'A' {
				next++
			}
			temp := 2*i + len(name) - next
			if move > temp {
				move = temp
			}
		}
	}
	return count + move
}
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
