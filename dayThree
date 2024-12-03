package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var Input []byte

var approachingSeq = []byte{
	'm', 'u', 'l', '(',
}

var allowingSeq = []byte{
	'd', 'o', '(', ')',
}

var disallowingSeq = []byte{
	'd', 'o', 'n', '\'', 't', '(', ')',
}

var sequencesSimple = [][]byte{
	approachingSeq,
}

var sequencesLogic = [][]byte{
	approachingSeq,
	disallowingSeq,
	allowingSeq,
}

var (
	canMultiply = true
	scanning    = false
	numCursor   = 0
	cursor      = 0
	sum         = 0
)

func isNum(b byte) bool {
	return b >= 48 && b <= 57
}

// need to switch out sequences simple for sequences logic (...) lazy
func getSequenceIndex(b byte) int {
	var skip = true
	var idx = -1
	for i, seq := range sequencesSimple {
		if cursor >= len(seq) || seq[cursor] != b {
			continue
		}
		if cursor == len(seq)-1 {
			idx = i
			break
		}
		skip = false
	}

	if skip {
		cursor = -1
	}

	cursor++

	return idx
}

func main() {
	var nums = make([]int, 2, 2)

	for _, b := range Input {
		var idx = getSequenceIndex(b)
		switch idx {
		case 0:
			scanning = true
			break
		case 1:
			canMultiply = false
			break
		case 2:
			canMultiply = true
			break
		}

		if idx >= 0 || !scanning {
			continue
		}

		if isNum(b) {
			nums[numCursor] = (nums[numCursor] * 10) + int(b-'0')
			continue
		}

		if b == ',' && numCursor == 0 {
			numCursor++
			continue
		}

		if b == ')' && canMultiply {
			sum += nums[0] * nums[1]
		}

		clear(nums)
		numCursor = 0
		scanning = false
	}

	fmt.Println(sum)
}
