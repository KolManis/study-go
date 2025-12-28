package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Get-Content input.txt | go run main.go

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

	isLocalMin := make([]bool, n)
	for i := 1; i < n-1; i++ {
		if a[i-1] > a[i] && a[i] < a[i+1] {
			isLocalMin[i] = true
		}
	}

	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if !isLocalMin[i] {
			result = append(result, a[i])
		}
	}

	// Вывод результата
	fmt.Println(len(result))
	for i, v := range result {
		fmt.Print(v)
		if i < len(result)-1 {
			fmt.Print(" ")
		}
	}
}
