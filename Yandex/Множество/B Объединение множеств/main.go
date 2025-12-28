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

	resultMap := make(map[int]bool)

	for i := 0; i < nInt; i++ {
		reqStr, _ := reader.ReadString('\n')
		reqSlicesStr := strings.Fields(strings.TrimSpace(reqStr))

		//req1, _ := strconv.Atoi(reqSlicesStr[0])

		for j := 1; j < len(reqSlicesStr); j++ {
			k, _ := strconv.Atoi(reqSlicesStr[j])
			resultMap[k] = true
		}
	}

	fmt.Println(len(resultMap))
}
