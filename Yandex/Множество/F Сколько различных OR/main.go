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

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	nums := strings.Split(line, " ")

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(nums[i])
	}

	allValues := make(map[int]bool)     // Все различные OR
	currentValues := make(map[int]bool) // OR для подмассивов, заканчивающихся на текущей позиции

	for i := 0; i < n; i++ {
		// Новый набор значений для позиции i
		newValues := make(map[int]bool)

		// Добавляем сам элемент
		newValues[arr[i]] = true
		allValues[arr[i]] = true

		// Объединяем со всеми значениями из предыдущего шага
		for val := range currentValues {
			newOr := val | arr[i]
			newValues[newOr] = true
			allValues[newOr] = true
		}

		// Обновляем currentValues для следующей итерации
		currentValues = newValues
	}

	fmt.Println(len(allValues))
}
