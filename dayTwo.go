package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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

func produceVariation(source []int, offset int) []int {
	if offset == -1 {
		return source
	}
	return slices.Delete(slices.Clone(source), offset, offset+1)
}

func variationsAsLoop(source []int, once bool) func(yield func([]int) bool) {
	return func(yield func([]int) bool) {
		for i := 0; i < len(source)+1; i++ {
			if !yield(produceVariation(source, i-1)) || once {
				return
			}
		}
	}
}

func calculateSafeSequences(values [][]int, tolerant bool) int {
	var safes = 0
	for _, split := range values {
		for sequence := range variationsAsLoop(split, !tolerant) {
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

	fmt.Println(calculateSafeSequences(numbers, false), "intolerant sequences")
	fmt.Println(calculateSafeSequences(numbers, true), "tolerant sequences")
}
