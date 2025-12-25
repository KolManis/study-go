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

	qStr, _ := reader.ReadString('\n')
	qint, _ := strconv.Atoi(strings.TrimSpace(qStr))

	list := make([]int, 0)

	for i := 0; i < qint; i++ {
		queryStr, _ := reader.ReadString('\n')
		querySliceStr := strings.Fields(strings.TrimSpace(queryStr))

		if len(querySliceStr) == 0 {
			continue
		}

		queryType, _ := strconv.Atoi(querySliceStr[0])

		switch queryType {
		case 1:
			x, _ := strconv.Atoi(querySliceStr[1])
			y, _ := strconv.Atoi(querySliceStr[2])

			if x == 0 {
				list = append([]int{y}, list...)
			} else {

				newList := make([]int, len(list)+1)
				copy(newList[:x], list[:x])
				newList[x] = y
				copy(newList[x+1:], list[x:])
				list = newList
			}

		case 2:
			x, _ := strconv.Atoi(querySliceStr[1])
			fmt.Println(list[x-1])

		case 3:
			x, _ := strconv.Atoi(querySliceStr[1])
			list = append(list[:x-1], list[x:]...)
		}
	}

}
