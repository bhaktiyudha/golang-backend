package main

import (
	"fmt"
	"strings"
)

func findFirstStringInBracket(str string) string {
	i := strings.IndexByte(str, '(')
	if i < 0 {
		return ""
	}
	i++
	j := strings.IndexByte(str, ')')
	if j < 0 {
		return ""
	}
	return str[i:j]
}

func main() {
	fmt.Println(findFirstStringInBracket("Hello (World)"))
}
