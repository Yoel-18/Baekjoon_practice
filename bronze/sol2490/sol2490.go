package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br = bufio.NewReader(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)
)

func main() {
	defer bw.Flush()
	for i := 0; i < 3; i++ {
		str, _ := br.ReadString('\n')
		count := 0
		for j := 0; j < len(str); j++ {
			if str[j] == '1' {
				count++
			}
		}
		if count == 1 {
			fmt.Fprintln(bw, "C")
		} else if count == 2 {
			fmt.Fprintln(bw, "B")
		} else if count == 3 {
			fmt.Fprintln(bw, "A")
		} else if count == 4 {
			fmt.Fprintln(bw, "E")
		} else {
			fmt.Fprintln(bw, "D")
		}
	}
}
