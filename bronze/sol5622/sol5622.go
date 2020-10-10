package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	count := 0
	for i := 0; i < len(str); i++ {
		s := str[i : i+1]
		switch s {
		case "A", "B", "C":
			count += 3

		case "D", "E", "F":
			count += 4

		case "G", "H", "I":
			count += 5

		case "J", "K", "L":
			count += 6

		case "M", "N", "O":
			count += 7

		case "P", "Q", "R", "S":
			count += 8

		case "T", "U", "V":
			count += 9

		case "W", "X", "Y", "Z":
			count += 10
		}
	}
	fmt.Print(count)
}
