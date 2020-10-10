package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	if sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
	}

	for i := 0; i < n; i++ {
		if sc.Scan() {
			a := strings.Fields(sc.Text())
			num, _ := strconv.Atoi(a[0])
			str := a[1]
			for _, v := range str {
				for j := 0; j < num; j++ {
					wr.WriteByte(byte(v))
				}

			}
			wr.WriteString("\n")
		}
	}
}
