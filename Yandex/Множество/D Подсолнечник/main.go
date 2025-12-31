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

	sets := make([]map[int]bool, n)

	for i := 0; i < n; i++ {
		reqStr, _ := reader.ReadString('\n')
		reqSlicesStr := strings.Fields(strings.TrimSpace(reqStr))

		sets[i] = make(map[int]bool)

		if len(reqSlicesStr) == 0 {
			continue
		}

		// k_i - размер множества (пропускаем)
		// fields[0] содержит k_i

		for j := 1; j < len(reqSlicesStr); j++ {
			num, _ := strconv.Atoi(reqSlicesStr[j])
			sets[i][num] = true
		}
	}

	core := make(map[int]bool)

	if n > 0 {
		for elem := range sets[0] {
			core[elem] = true
		}

		// Последовательно пересекаем с остальными множествами
		for i := 1; i < n; i++ {
			newCore := make(map[int]bool)
			for elem := range core {
				if sets[i][elem] {
					newCore[elem] = true
				}
			}
			core = newCore
		}
	}

	coreSlice := make([]int, 0, len(core))
	for elem := range core {
		coreSlice = append(coreSlice, elem)
	}

	// fmt.Println("Ядро C:", coreSlice)
	// fmt.Println("Размер ядра |C|:", len(core))

	// 2. Проверяем условие подсолнечника
	elementToSet := make(map[int]int)
	isValid := true

	for i := 0; i < n && isValid; i++ {
		for elem := range sets[i] {
			if !core[elem] {
				if prevSet, exists := elementToSet[elem]; exists {
					if prevSet != i {
						// fmt.Println("NO - элементы лепестков пересекаются")
						// fmt.Printf("Элемент %d есть в множествах %d и %d\n", elem, prevSet+1, i+1)
						isValid = false
						break
					}
				} else {
					elementToSet[elem] = i
				}
			}
		}
	}

	if !isValid {
		fmt.Println("NO")
		return
	}

	// 3. Вычисляем размеры лепестков
	fmt.Println("YES")
	fmt.Println(len(core))

	for i := 0; i < n; i++ {
		// Размер лепестка = размер множества - размер ядра
		petalSize := len(sets[i]) - len(core)
		fmt.Print(petalSize)
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
