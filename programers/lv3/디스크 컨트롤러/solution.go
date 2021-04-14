package main

import "sort"

func main() {

}
func solution(jobs [][]int) int {

	sort.Slice(jobs, func(i, k int) bool {
		if jobs[i][1] != jobs[k][1] {
			return jobs[i][1] < jobs[k][1]
		} else {
			return jobs[i][0] < jobs[k][0]
		}
	})

	ts := 0
	count := len(jobs)
	sec := 0

	for len(jobs) > 0 {
		found := false

		for i, job := range jobs {
			if job[0] <= sec {
				sec += job[1]
				ts += sec - job[0]

				copy(jobs[i:], jobs[i+1:])
				jobs = jobs[:len(jobs)-1]

				found = true
				break
			}
		}

		if !found {
			sec++
		}
	}

	return ts / count
}
