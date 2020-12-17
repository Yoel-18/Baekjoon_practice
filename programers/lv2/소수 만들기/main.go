package main

func main() {

}

func solution(nums []int) int {
	answer := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				sum = nums[i] + nums[j] + nums[k]
				if isPrime(sum) {
					answer++
				}
			}
		}
	}
	return answer
}

func isPrime(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
