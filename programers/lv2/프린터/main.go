package main

//	백준이랑 같은 문제인듯?

func main() {

}
func solution(priorities []int, location int) int {
	q := make([]pri, 0)
	count := 0
	for j := 0; j < len(priorities); j++ {
		q = append(q, pri{j, priorities[j]})
	}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		check := true

		for _, que := range q {
			if que.priority > cur.priority {
				check = false
			}
		}
		if check {
			count++
			if cur.index == location {
				break
			}
		} else {
			q = append(q, cur)
		}
	}
	return count
}

type pri struct {
	index, priority int
}
