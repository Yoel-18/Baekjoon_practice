package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func formingMagicSquare(s [][]int32) int32 {
	ans := [][][]int{
		{{8, 1, 6},
			{3, 5, 7},
			{4, 9, 2}},

		{{6, 1, 8},
			{7, 5, 3},
			{2, 9, 4}},

		{{4, 9, 2},
			{3, 5, 7},
			{8, 1, 6}},

		{{2, 9, 4},
			{7, 5, 3},
			{6, 1, 8}},

		{{8, 3, 4},
			{1, 5, 9},
			{6, 7, 2}},

		{{4, 3, 8},
			{9, 5, 1},
			{2, 7, 6}},

		{{6, 7, 2},
			{1, 5, 9},
			{8, 3, 4}},

		{{2, 7, 6},
			{9, 5, 1},
			{4, 3, 8}}}

	min_cost := 100
	calc := 0

	for n := 0; n < 8; n++ {
		calc = 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				calc += abs(int(s[i][j]) - ans[n][i][j])
			}
		}
		min_cost = min(calc, min_cost)
	}

	return int32(min_cost)
}
func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(readLine(reader), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
