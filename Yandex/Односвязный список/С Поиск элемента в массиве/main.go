package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	nq := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(nq[0])
	q, _ := strconv.Atoi(nq[1])

	scanner.Scan()
	numbers := strings.Fields(scanner.Text())

	const MAX_VAL = 100001
	firstIndex := make([]int, MAX_VAL)

	for i := range firstIndex {
		firstIndex[i] = -1
	}

	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(numbers[i])
		if firstIndex[num] == -1 {
			firstIndex[num] = i + 1
		}
	}

	for i := 0; i < q; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())

		idx := firstIndex[x]
		fmt.Println(idx)
	}
}
