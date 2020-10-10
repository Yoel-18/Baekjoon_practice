package main

func main() {

}
func sum(a []int) int {
	count := 0
	for i := 0; i < len(a); i++ {
		count += a[i]
	}
	return count

	// for _, a := range a {
	// 	r += a
	// }
}
