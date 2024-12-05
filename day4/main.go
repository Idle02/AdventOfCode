package main

import (
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var Input []byte


var (
	xmas       = []byte(`XMAS`)
	crlf       = []byte("\r\n")
	mas        = []byte(`MAS`)
	sam        = []byte(`SAM`)
)

func traverseInDirection(data [][]byte, x, y, dx, dy, bound int) int {
	cx, cy := x, y
	// skip the x
	for i := 1; i < 4; i++ {
		cx, cy = cx+dx, cy+dy

		if cx < 0 || cx > bound || cy < 0 || cy > bound {
			return 0
		}

		if data[cy][cx] != xmas[i] {
			return 0
		}
	}
	return 1
}

func directionsAsLoop() func(yield func(int, int) bool) {
	return func(yield func(int, int) bool) {
		// muffin approach
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if !yield(dx, dy) {
					return
				}
			}
		}
	}
}

func countXMAS(data [][]byte, bounds int) int {
	var xmases = 0
	for y := 0; y <= bounds; y++ {
		for x := 0; x <= bounds; x++ {
			if data[y][x] != 'X' {
				continue
			}
			for dx, dy := range directionsAsLoop() {
				xmases += traverseInDirection(data, x, y, dx, dy, bounds)
			}
		}
	}
	return xmases
}

func checkX(data [][]byte, cx, cy, bounds int) int {
	if cx == 0 || cy == 0 || cx >= bounds || cy >= bounds {
		return 0
	}
	var diag1 = []byte{data[cy-1][cx-1], data[cy][cx], data[cy+1][cx+1]}
	var diag2 = []byte{data[cy-1][cx+1], data[cy][cx], data[cy+1][cx-1]}
	if (!bytes.Equal(diag1, mas) && !bytes.Equal(diag1, sam)) || (!bytes.Equal(diag2, mas) && !bytes.Equal(diag2, sam)) {
		return 0
	}
	return 1
}

func countMAS_X(data [][]byte, bounds int) int {
	var xmases = 0
	for y := 0; y <= bounds; y++ {
		for x := 0; x <= bounds; x++ {
			if data[y][x] != 'A' {
				continue
			}
			xmases += checkX(data, x, y, bounds)
		}
	}
	return xmases
}

func main() {
	var data = bytes.Split(Input, crlf)

	var wrapLength = len(data) - 1

	fmt.Println(countXMAS(data, wrapLength))
	fmt.Println(countMAS_X(data, wrapLength))
}
