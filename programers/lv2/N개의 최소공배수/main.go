package main

func main() {

}
func solution(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	answer := 0
	for i := 1; i < len(arr); i++ {
		answer = lcm(arr[i], arr[i-1])
		arr[i] = answer
	}
	return answer
}

//	최대공약수
func gcd(a, b int) int {
	for b != 0 {
		r := a % b
		a = b
		b = r
	}
	return a
}

//	최소공배수
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
