package main

import "fmt"

var (
	n   int
	m   int
	num [100]int
)

func main() {
	fmt.Scan(&n, &m)
	for i := 0; i < n; i++ {
		fmt.Scan(&num[i])
	}
	fmt.Print(find())
}

func find() int {
	result := 0
	for i := 0; i < n-2; i++ {
		if num[i] > m {
			continue
		}
		for j := i + 1; j < n-1; j++ {
			if num[i]+num[j] > m {
				continue
			}
			for k := j + 1; k < n; k++ {
				sum := num[i] + num[j] + num[k]
				if m == sum {
					return sum
				}
				if result < sum && sum < m {
					result = sum
				}
			}
		}
	}
	return result
}
