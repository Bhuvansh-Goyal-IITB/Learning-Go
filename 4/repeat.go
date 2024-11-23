package main

func Repeat(char string, repeatNum int) string {
	var repeat string
	for i := 0; i < repeatNum; i++ {
		repeat += char
	}
	return repeat
}
