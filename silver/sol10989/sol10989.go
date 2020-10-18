package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br   = bufio.NewReader(os.Stdin)
	bw   = bufio.NewWriter(os.Stdout)
	n    int
	nums [10001]int
)

//시간초과...???
func main() {
	defer bw.Flush()
	fmt.Fscan(br, &n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(br, &num)
		nums[num]++
	}
	for i := 1; i <= 10000; i++ {
		if nums[i] > 0 {
			for j := 0; j < nums[i]; j++ {
				fmt.Println(i)
			}
		}
	}
}
