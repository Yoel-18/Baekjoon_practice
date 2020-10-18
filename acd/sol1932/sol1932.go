package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// var (
// 	br  = bufio.NewReader(os.Stdin)
// 	bw  = bufio.NewWriter(os.Stdout)
// 	n   int
// 	t   [500][500]int
// 	max = 0
// )
//원래풀이
// func main() {
// 	defer bw.Flush()
// 	fmt.Fscan(br, &n)
// 	fmt.Fscan(br, &t[0][0])
// 	for i := 1; i < n; i++ {
// 		for j := 0; j <= i; j++ {
// 			fmt.Fscan(br, &t[i][j])
// 			if j == 0 {
// 				t[i][j] += t[i-1][j]
// 			} else if j == i {
// 				t[i][j] += t[i-1][j-1]
// 			} else {
// 				t[i][j] += maxInt(t[i-1][j-1], t[i-1][j])
// 			}
// 			max = maxInt(max, t[i][j])
// 		}
// 	}
// 	fmt.Fprint(bw, max)
// }
// func maxInt(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
