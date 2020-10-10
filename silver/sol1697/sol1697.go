package main

import "fmt"

var (
	n, k  int
	q     []int
	check [100001]int
)

func main() {
	fmt.Scan(&n, &k)
	findByBFS()
	fmt.Print(check[k])
}
func findByBFS() {
	q = append(q, n)
	for len(q) > 0 {
		x := q[0]
		q = q[1:len(q)]
		if x == k {
			break
		}
		if x > 0 && check[x-1] == 0 {
			q = append(q, x-1)
			check[x-1] = check[x] + 1
		}
		if x < 100000 && check[x+1] == 0 {
			q = append(q, x+1)
			check[x+1] = check[x] + 1
		}
		if x*2 <= 100000 && check[x*2] == 0 {
			q = append(q, x*2)
			check[x*2] = check[x] + 1
		}
	}
}
