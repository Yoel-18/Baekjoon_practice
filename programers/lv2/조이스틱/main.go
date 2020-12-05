package main

func main() {

}
func solution(name string) int {
	answer := 0
	for i := 0; i < len(name); i++ {
		if name[i] != 'A' {
			up := (int)(name[i] - 'A')
			down := (int)(1 + 'Z' - name[i])
			if up > down {
				answer += down
			} else {
				answer += up
			}
		}
	}

	minMove := len(name) - 1
	for i := 0; i < len(name); i++ {
		if name[i] != 'A' {
			next := i + 1
			for next < len(name) && name[next] == 'A' {
				next++
			}
			temp := 2*i + len(name) - next
			if minMove > temp {
				minMove = temp
			}
		}
	}
	return answer + minMove
}
