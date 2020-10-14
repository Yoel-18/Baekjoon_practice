package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var rr *bufio.Reader = bufio.NewReader(os.Stdin)
var ww *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(ww, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(rr, f, a...) }

//와우...풀이 방법 기가 맥힌다...
func main() {
	defer ww.Flush()

	var s string
	scanf("%s\n", &s)

	s = strings.Replace(s, "c=", "1", -1)
	s = strings.Replace(s, "c-", "2", -1)
	s = strings.Replace(s, "dz=", "3", -1)
	s = strings.Replace(s, "d-", "4", -1)
	s = strings.Replace(s, "lj", "5", -1)
	s = strings.Replace(s, "nj", "6", -1)
	s = strings.Replace(s, "s=", "7", -1)
	s = strings.Replace(s, "z=", "8", -1)

	printf("%d", len(s))
}
