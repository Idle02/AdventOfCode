package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func isValidSequence(sequence []int) bool {
	ascending := sequence[1] > sequence[0]
	for i := 0; i < len(sequence)-1; i++ {
		var delta = sequence[i+1] - sequence[i]
		if (delta > 0 != ascending) || delta == 0 || abs(delta) > 3 {
			return false
		}
	}
	return true
}

func calculateSafesIntolerant(values [][]int) int {
	var safes = 0
	for _, split := range values {
		if isValidSequence(split) {
			safes++
		}
	}
	return safes
}

func produceVariation(source []int, offset int) []int {
	if offset == -1 {
		return source
	}
	var duplicate = make([]int, len(source))
	copy(duplicate, source)
	return append(duplicate[:offset], duplicate[offset+1:]...)
}

func variationsAsLoop(source []int) func(yield func([]int) bool) {
	return func(yield func([]int) bool) {
		for i := 0; i < len(source)+1; i++ {
			if !yield(produceVariation(source, i-1)) {
				return
			}
		}
	}
}

func calculateSafesTolerant(values [][]int) int {
	var safes = 0

	for _, split := range values {
		for sequence := range variationsAsLoop(split) {
			if isValidSequence(sequence) {
				safes++
				break
			}
		}
	}
	return safes
}

func main() {
	f, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatalln(err)
	}

	var split = strings.Split(string(f), "\r\n")

	var numbers = make([][]int, len(split))

	for i, seq := range split {
		numbers[i] = func(s string) []int {
			var splits = strings.Split(s, " ")
			var out = make([]int, len(splits))
			for j, num := range splits {
				out[j], _ = strconv.Atoi(num)
			}
			return out
		}(seq)
	}

	fmt.Println(calculateSafesIntolerant(numbers), "intolerant sequences")
	fmt.Println(calculateSafesTolerant(numbers), "tolerant sequences")
}
