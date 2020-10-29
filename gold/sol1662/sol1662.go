package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br = bufio.NewReader(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)
	// s     string
	// str   []string
	// st    []int
	// paren [50]int
	s []byte
)

func main() {
	fmt.Scan(&s)
	fmt.Println(find())
}
func find() int {
	r := 0
	for {
		switch {
		case len(s) == 0: //	종료
			return r
		case s[0] == ')': //	닫히면 pop
			s = s[1:]
			return r
		case len(s) > 1 && s[1] == '(': //	열리는 괄호면 곱하는 수 찾고, push
			k := int(s[0]) - '0'
			s = s[2:]
			r += k * find()
		default:
			s = s[1:]
			r++
		}
	}
}

// func main() {
// 	fmt.Fscan(br, &s)
// 	str = strings.Split(s, "")
// 	// 9가 1번 반복 -> 79
// 	// 79가 두번 반복 -> 567979
// 	// 567979가 3번 반복 -> 3567979567979567979
// 	//	123(123) -> 12 123123123
// 	for i := 0; i < len(str); i++ {
// 		if str[i] == "(" {
// 			st = append(st, i)
// 		}
// 		if str[i] == ")" {
// 			cur := st[len(st)-1]
// 			st = st[:len(st)]
// 			paren[cur] = i
// 		}
// 	}
// 	fmt.Print(find(0, len(str)))
// }
// func find(start, end int) int {
// 	sLen := 0
// 	for i := start; i < end; i++ {
// 		if str[i] == "(" {
// 			sLen += (str[i-1]-"0")*find(i+1, paren[i]) - 1
// 			i = paren[i]
// 		} else {
// 			sLen++
// 		}
// 	}
// 	return sLen
// }
