package main

import "strings"

func Repeat(char string, repeatNum int) string {
	var repeat strings.Builder
	for i := 0; i < repeatNum; i++ {
		repeat.WriteString(char)
	}
	return repeat.String()
}
