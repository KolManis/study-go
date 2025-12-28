package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	arrStr, _ := reader.ReadString('\n')
	parts := strings.Fields(strings.TrimSpace(arrStr))

	arr := make([]int, n)
	for i := 0; i < n && i < len(parts); i++ {
		num, _ := strconv.Atoi(parts[i])
		arr[i] = num
	}

	minDiff := 1 << 30
	minI, minJ := 0, 1

	maxDiff := -1 << 30
	maxI, maxJ := 0, 1

	leftMinIdx := 0
	leftMaxIdx := 0

	for j := 1; j < n; j++ {
		if arr[j-1] < arr[leftMinIdx] {
			leftMinIdx = j - 1
		}
		if arr[j-1] > arr[leftMaxIdx] {
			leftMaxIdx = j - 1
		}

		diffMin := arr[leftMinIdx] - arr[j]
		if diffMin < minDiff {
			minDiff = diffMin
			minI, minJ = leftMinIdx, j
		} else if diffMin == minDiff {
			if leftMinIdx < minI || (leftMinIdx == minI && j < minJ) {
				minI, minJ = leftMinIdx, j
			}
		}

		diffMax := arr[leftMaxIdx] - arr[j]
		if diffMax > maxDiff {
			maxDiff = diffMax
			maxI, maxJ = leftMaxIdx, j
		} else if diffMax == maxDiff {
			if leftMaxIdx < maxI || (leftMaxIdx == maxI && j < maxJ) {
				maxI, maxJ = leftMaxIdx, j
			}
		}
	}

	// Вывод (индексы с 1)
	fmt.Printf("%d %d\n", minI+1, minJ+1)
	fmt.Printf("%d %d\n", maxI+1, maxJ+1)
}
