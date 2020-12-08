package main

func main() {

}

var (
	a [31]int
	v [31]int
)

func solution(clothes [][]string) int {
	answer := 0
	size := len(clothes)
	ret := 1
	for i := 0; i < 31; i++ {
		v[i] = 0
	}
	for i := 0; i < size; i++ {
		if v[i] == 1 {
			continue
		}
		for j := i; j < size; j++ {
			if clothes[i][1] == clothes[j][1] {
				a[i]++
				v[j]++
			}
		}
	}
	for i := 0; i < 31; i++ {
		if a[i] > 0 {
			ret *= (a[i] + 1)
		}
	}
	answer = ret - 1
	return answer
}

// func solution(clothes [][]string) int {
//     counter := make(map[string]int)

//     for _, cloth := range clothes {
//         counter[cloth[1]]++
//     }

//     x := 1
//     for _, count := range counter {
//         x *= count + 1
//     }
//     x -= 1

//     return x
// }
