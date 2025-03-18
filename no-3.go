package main

import (
	"fmt"
	"strings"
)

func isBalanced(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	stack := []rune{}

	for i := 0; i < len(s); i++ {
		ch := rune(s[i])
		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if (ch == ')' && top != '(') ||
				(ch == '}' && top != '{') ||
				(ch == ']' && top != '[') {
				return false
			}
		}
	}
	return len(stack) == 0 // true kalo ampe akhir stack kosong
}

func main() {
	var input string

	fmt.Scanf("%[^\n]", &input) // baca include space sampe newline

	if isBalanced(input) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
