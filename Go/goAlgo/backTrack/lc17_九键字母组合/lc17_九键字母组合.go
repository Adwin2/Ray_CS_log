package main

import "fmt"

type str string

var letterMap = [10]string{
	"",
	"",
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

var res []string

func backTrack(s, digits string, index int) {
	if index == len(digits) {
		res = append(res, s)
		return
	}
	digit := digits[index] - '0'
	letters := letterMap[digit]
	for i := 0; i < len(letters); i++ {
		s += string(letters[i])
		// fmt.Println(s)
		backTrack(s, digits, index+1)
		s = s[:len(s)-1]
	}
}

func main() {
	backTrack("", "23", 0)
	fmt.Println(res)
}
