package main

func btoi(b []byte) int {
	n := 0
	for _, c := range b {
		n = n*10 + int(c-'0')
	}
	return n
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func getDistance(data []byte) int {
	var numbersRight = make([]int, 1000)
	var numbersLeft = make([]int, 1000)

	for i, line := range bytes.Split(data, []byte("\r\n")) {
		var split = bytes.Split(line, []byte("   "))
		numbersRight[i] = btoi(split[1])
		numbersLeft[i] = btoi(split[0])
	}

	slices.Sort(numbersRight)
	slices.Sort(numbersLeft)

	var deltas = 0

	for i := 0; i < 1000; i++ {
		deltas += abs(numbersRight[i] - numbersLeft[i])
	}

	return deltas
}

func getSimilarity(data []byte) int {
	var numberList = make(map[int]int, 1000)
	var numbersLeft = make([]int, 1000)

	for i, line := range bytes.Split(data, []byte("\r\n")) {
		var split = bytes.Split(line, []byte("   "))
		numbersLeft[i] = btoi(split[0])
		numberList[btoi(split[1])]++
	}

	var similarity = 0

	for i := 0; i < 1000; i++ {
		similarity += numbersLeft[i] * numberList[numbersLeft[i]]
	}

	return similarity
}

func main() {
	f, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("distance: ", getDistance(f))
	fmt.Println("similarity: ", getSimilarity(f))
}
