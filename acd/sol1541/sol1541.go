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

//	곱셈과 나눗셈이 있다고 가정
//	문제가 쉽지 않네용
