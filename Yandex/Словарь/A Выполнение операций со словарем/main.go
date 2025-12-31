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

	nString, _ := reader.ReadString('\n')
	nInt, _ := strconv.Atoi(strings.TrimSpace(nString))

	ResultMap := make(map[int]int)
	answer := make([]int, 0, 0)
	for i := 0; i < nInt; i++ {
		qString, _ := reader.ReadString('\n')
		qSlicesString := strings.Fields(strings.TrimSpace(qString))

		if qSlicesString[0] == "1" {
			x, _ := strconv.Atoi(qSlicesString[1])
			y, _ := strconv.Atoi(qSlicesString[2])
			ResultMap[x] = y
		} else if qSlicesString[0] == "2" {
			x, _ := strconv.Atoi(qSlicesString[1])
			if v, ok := ResultMap[x]; !ok {
				answer = append(answer, -1)
			} else {
				answer = append(answer, v)
			}
		}
	}

	for _, v := range answer {
		fmt.Println(v)
	}
}
