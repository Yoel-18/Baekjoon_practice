package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReaderSize(os.Stdin, 1000000)
	input, _, _ := br.ReadLine()
	str := bytes.TrimSpace([]byte(input))

	if len(str) == 0 {
		fmt.Print(0)
		return
	}
	count := 1
	target := byte(' ')
	for _, char := range str {
		if char == target {
			count++
		}
	}
	fmt.Print(count)
}
