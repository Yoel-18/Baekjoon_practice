package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type member struct {
	age  int
	name string
}

func main() {
	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)

	n := 0
	fmt.Fscanln(br, &n)

	memberList := make([]member, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(br, &memberList[i].age, &memberList[i].name)
	}

	//	익명함수 사용
	sort.SliceStable(memberList, func(i, j int) bool {
		return memberList[i].age < memberList[j].age
	})

	for i := 0; i < n; i++ {
		fmt.Fprintln(bw, memberList[i].age, memberList[i].name)
	}

	bw.Flush()
}
