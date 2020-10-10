package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

const (
	maxLen = 1000
)

var dp [maxLen + 1][maxLen + 1]int

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var s1, s2 string
	fmt.Fscan(reader, &s1, &s2)

	s1n, s2n := len(s1), len(s2)
	ar1, ar2 := []rune(s1), []rune(s2)

	for i := 1; i <= s1n; i++ {
		for j := 1; j <= s2n; j++ {
			if ar1[i-1] == ar2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	fmt.Fprintln(writer, dp[s1n][s2n])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
