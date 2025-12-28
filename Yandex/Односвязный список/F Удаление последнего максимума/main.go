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

	numsStr, _ := reader.ReadString('\n')
	numSlicesStr := strings.Fields(strings.TrimSpace(numsStr))

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(numSlicesStr[i])
	}

	currentMin := a[0]
	currentMinI := 0
	for i := 0; i < n; i++ {
		if a[i] >= currentMin {
			currentMin = a[i]
			currentMinI = i
		}
	}

	for i := 0; i < n; i++ {
		if i == currentMinI {
			continue
		}
		fmt.Print(a[i])
		if i < len(a)-1 {
			fmt.Print(" ")
		}
	}
}
