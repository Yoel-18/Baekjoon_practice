package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	br     = bufio.NewReader(os.Stdin)
	bw     = bufio.NewWriter(os.Stdout)
	input  string
	answer int
)

//	나눗셈, 곱셈 추가
//	사칙연산 순서를 지킨다고 가정
//	경우의 수가 너무 많습니다...
func main() {
	defer bw.Flush()
	fmt.Fscan(br, &input)
	str := strings.Split(input, "-")
	for i := 0; i < len(str); i++ {
		p := strings.Split(str[i], "+")
		sum := 0
		for _, number := range p {
			num, _ := strconv.Atoi(number)
			sum += num
		}
		if i == 0 {
			answer += sum
		} else {
			answer -= sum
		}
	}
	fmt.Fprint(bw, answer)
}
