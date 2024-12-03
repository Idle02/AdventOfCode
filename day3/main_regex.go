package main

import (
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var Input string

var numPattern = regexp.MustCompile(`([0-9]){1,3}`)

func strToSum(str string) int {
	var split = numPattern.FindAllString(str, 2)
	var n1, _ = strconv.Atoi(split[0])
	var n2, _ = strconv.Atoi(split[1])
	return n1 * n2
}

func simple() int {
	mulPattern := regexp.MustCompile(`mul\(([0-9]{1,3})+,([0-9]{1,3})+\)`)

	var sum = 0
	for _, str := range mulPattern.FindAllString(Input, -1) {
		sum += strToSum(str)
	}

	return sum
}

func logic() int {
	logicPattern := regexp.MustCompile(`(mul\(([0-9]{1,3})+,([0-9]{1,3})+\))|(do\(\))|(don't\(\))`)

	var canMult = true
	var sum = 0

	for _, str := range logicPattern.FindAllString(Input, -1) {
		switch str {
		case "do()":
			canMult = true
			break
		case "don't()":
			canMult = false
			break
		default:
			if !canMult {
				break
			}
			sum += strToSum(str)
		}
	}

	return sum
}

func main() {
	fmt.Println(simple())
	fmt.Println(logic())
}
