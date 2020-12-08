package main

func main() {
	solution([]int{95, 90, 99, 99, 80, 99}, []int{1, 1, 1, 1, 1, 1})
}
func solution(progresses []int, speeds []int) []int {
	st := make([]int, 0)
	count := 1
	cur := (100 - progresses[0]) / speeds[0]
	if (100-progresses[0])%speeds[0] != 0 {
		cur++
	}
	for i := 1; i < len(progresses); i++ {
		next := (100 - progresses[i]) / speeds[i]
		if (100-progresses[i])%speeds[i] != 0 {
			next++
		}
		if cur >= next {
			count++
		} else {
			st = append(st, count)
			count = 1
			cur = next
		}
		if i == len(progresses)-1 {
			st = append(st, count)
		}
	}
	return st
}
