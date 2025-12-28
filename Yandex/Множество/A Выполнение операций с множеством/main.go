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

		req1, _ := strconv.Atoi(reqSlicesStr[0])
		req2, _ := strconv.Atoi(reqSlicesStr[1])
		if req1 == 2 {
			if _, ok := resultMap[req2]; !ok {
				fmt.Println(0)
			} else {
				fmt.Println(1)
			}
		} else {
			resultMap[req2] = true
		}
	}
}
