package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	br  = bufio.NewReader(os.Stdin)
	bw  = bufio.NewWriter(os.Stdout)
	n   int
	aa  []int
	op  = [4]int{}
	max = -1 * math.MaxInt64
	min = math.MaxInt64
)

func main() {
	defer bw.Flush()
	fmt.Fscan(br, &n)
	aa = make([]int, n)

	for i := range aa {
		fmt.Fscan(br, &aa[i])
	}
	for i := range op {
		fmt.Fscan(br, &op[i])
	}

	find(aa[0], 1, op[0], op[1], op[2], op[3])

	fmt.Fprintln(bw, max)
	fmt.Fprint(bw, min)
}

func find(result, index, a, s, m, d int) {
	if index == n {
		if result > max {
			max = result
		}
		if min > result {
			min = result
		}
		return
	}

	//	덧셈
	if a != 0 {
		find(result+aa[index], index+1, a-1, s, m, d)
	}

	//	뺄셈
	if s != 0 {
		find(result-aa[index], index+1, a, s-1, m, d)
	}

	//	곱셈
	if m != 0 {
		find(result*aa[index], index+1, a, s, m-1, d)
	}

	//	나눗셈
	if d != 0 {
		if result < 0 {
			result *= -1
			result /= aa[index]
			find(-1*result, index+1, a, s, m, d-1)
		} else {
			find(result/aa[index], index+1, a, s, m, d-1)
		}
	}
}
