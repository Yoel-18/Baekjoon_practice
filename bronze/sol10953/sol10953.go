package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	defer bw.Flush()
	t := 0
	fmt.Fscan(br, &t)
	for i := 0; i < t; i++ {
		var str string
		fmt.Fscan(br, &str)
		num := strings.Split(str, ",")
		a, _ := strconv.Atoi(num[0])
		b, _ := strconv.Atoi(num[1])
		fmt.Fprintln(bw, (a + b))
	}
}
