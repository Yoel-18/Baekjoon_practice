package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc              = bufio.NewScanner(os.Stdin)
	n, m            int
	h               = [50][50]int{}
	chicken, person = make([]point, 0), make([]point, 0)
	selChicken      []point
	result          = 1 << 30
)

//	치킨집 어렵네용...
func main() {
	n, m = nextInt(), nextInt()
	selChicken = make([]point, m)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			h[i][j] = nextInt()
			if h[i][j] == 1 { //	집 위치 추가
				person = append(person, point{i, j})
			} else if h[i][j] == 2 { //	치킨 위치 추가
				chicken = append(chicken, point{i, j})
			}
		}
	}
	find(0, 0)
	fmt.Print(result)

}

//	cs: 선택된 치킨집 수
func find(cs, start int) {
	//	선택된 치킨집이 m과 같으면 계산
	if cs == m {
		sum := 0
		for _, h := range person {
			minDist := 1 << 30
			//	선택된 치킨집과 각 집과의 최소거리를 구해서 합을 구한다
			for i := 0; i < m; i++ {
				//	선택된 치킨집과 집과의 거리
				d := dis(selChicken[i], h)
				//	이전에 선택된 치킨집과의 거리와 비교해서 최소값 저장
				minDist = minInt(minDist, d)
			}
			//	거리의 최솟값을 합에 더함
			sum += minDist
		}
		//	반복문 종료 후 거리 총합 비교 후 종료
		result = minInt(result, sum)
		return
	}
	//	차례대로 치킨집 집어넣음 ex) 0,1,2...0,1,3...0,1,4
	for i := start; i < len(chicken); i++ {
		selChicken[cs] = chicken[i]
		find(cs+1, i+1)
	}
}
func dis(d1, d2 point) int {
	return abs(d1.x-d2.x) + abs(d1.y-d2.y)
}

type point struct {
	x, y int
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
