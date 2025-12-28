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
	words := strings.Fields(strings.TrimSpace(numsStr))

	currentMin := 1 << 30
	var result strings.Builder

	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(words[i])
		if num < currentMin {
			currentMin = num
		}

		if i > 0 {
			result.WriteString(" ")
		}
		result.WriteString(strconv.Itoa(currentMin))
	}

	fmt.Println(result.String())
}
