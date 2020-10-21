package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	defer bw.Flush()
	var str string
	fmt.Fscan(br, &str)
	s := strings.Split(str, "")
	st := make([]int, 0)
	answer := 0
	for i := 0; i < len(s); i++ {
		if s[i] == "(" {
			st = append(st, i)
		} else {
			if st[len(st)-1] == (i - 1) {
				st = st[:len(st)-1]
				answer += len(st)
			} else {
				st = st[:len(st)-1]
				answer++
			}

		}
	}
	fmt.Fprint(bw, answer)
}
