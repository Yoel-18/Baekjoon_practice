package main

import "fmt"

var res = -1

func main() {
	var n, m int
	var vmap [][]int
	fmt.Scan(&n, &m)
	vmap = make([][]int, n)
	for i := 0; i < n; i++ {
		vmap[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&vmap[i][j])
		}
	}
	search(vmap, 0, 0, 0)
	fmt.Print(res)
}

//d: 벽의 갯수
func search(vmap [][]int, d, r, c int) {
	if d == 3 {
		var virus []pos //	바이러스담기
		count := 0
		for i := range vmap {
			for j := range vmap[i] {
				if vmap[i][j] == 2 {
					var tmp pos
					tmp.r = i
					tmp.c = j
					vmap[i][j] = 0
					virus = append(virus, tmp)
				}
			}
		}
		for _, v := range virus {
			infectVirus(vmap, v.r, v.c)
		}
		for i := range vmap {
			for j := range vmap[i] {
				if vmap[i][j] == 0 {
					count++
				}
			}
		}
		if res < count {
			res = count
		}
		return
	}
	for i := r; i < len(vmap); i++ {
		var j int
		if i == r {
			j = c
		} else {
			j = 0
		}
		for ; j < len(vmap[i]); j++ {
			newVmap := make([][]int, len(vmap))
			for ii := range newVmap {
				newVmap[ii] = append([]int{}, vmap[ii]...)
			}
			if newVmap[i][j] == 0 {
				newVmap[i][j] = 1
				search(newVmap, d+1, i, j)
			}
		}
	}
}

//바이러스 전파
func infectVirus(vmap [][]int, r, c int) {
	if vmap[r][c] == 0 {
		vmap[r][c] = 2
		if r+1 < len(vmap) {
			infectVirus(vmap, r+1, c)
		}
		if c+1 < len(vmap[r]) {
			infectVirus(vmap, r, c+1)
		}
		if r-1 >= 0 {
			infectVirus(vmap, r-1, c)
		}
		if c-1 >= 0 {
			infectVirus(vmap, r, c-1)
		}
	}
}

type pos struct {
	r int
	c int
}
