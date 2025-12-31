package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Get-Content input.txt | go run main.go

func main() {
	reader := bufio.NewReader(os.Stdin)

	nStr, _ := reader.ReadString('\n')
	nInt, _ := strconv.Atoi(strings.TrimSpace(nStr))

	resultMap := make(map[int]int)
	count := 0
	for i := 0; i < nInt; i++ {
		reqStr, _ := reader.ReadString('\n')
		reqSlicesStr := strings.Fields(strings.TrimSpace(reqStr))

		if i == 0 {
			for j := 1; j < len(reqSlicesStr); j++ {
				num, _ := strconv.Atoi(reqSlicesStr[j])
				resultMap[num] = 1
			}
			continue
		}

		for j := 1; j < len(reqSlicesStr); j++ {
			k, _ := strconv.Atoi(reqSlicesStr[j])
			if _, ok := resultMap[k]; ok {
				resultMap[k]++
			}
		}
	}

	for _, k := range resultMap {
		if k == nInt {
			count++
		}
	}

	fmt.Println(count)
}
