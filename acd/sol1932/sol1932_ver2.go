package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br  = bufio.NewReader(os.Stdin)
	bw  = bufio.NewWriter(os.Stdout)
	n   int
	num [500][500]int
	t   [500][500]int
	max = 0
)

//중간과정 검사
func main() {
	defer bw.Flush()
	fmt.Fscan(br, &n)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Fscan(br, &t[i][j])
			num[i][j] = t[i][j]
		}
	}
	fmt.Fprint(bw, "처음 입력받은 삼각형\n\n")
	printTri(num)
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				t[i][j] += t[i-1][j]
			} else if j == i {
				t[i][j] += t[i-1][j-1]
			} else {
				t[i][j] += maxInt(t[i-1][j-1], t[i-1][j])
			}
			num[i][j] = t[i][j]
			max = maxInt(max, t[i][j])
			fmt.Fprintf(bw, "\n\n%d번 줄, %d번 칸 비교 후 변경\n", i+1, j+1)
			fmt.Fprint(bw, "현재 최댓값: ", max, "\n")
			printTri(num)
		}
	}
	fmt.Fprint(bw, "\n최댓값: ", max, "\n")
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func printTri(tri [500][500]int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= 3*(n-i)/2; j++ {
			fmt.Fprint(bw, " ")
		}
		for j := 0; j < i+1; j++ {
			fmt.Fprintf(bw, "%3d", tri[i][j])
		}
		fmt.Fprintln(bw)
	}
}
