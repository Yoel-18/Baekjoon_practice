package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	defer bw.Flush()
	var str string
	var err error
	for {
		str, err = br.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprint(bw, str)
	}
}
