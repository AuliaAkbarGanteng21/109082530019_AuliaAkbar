package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func selectionSortAscending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func selectionSortDescending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[idxMax] {
				idxMax = j
			}
		}
		arr[i], arr[idxMax] = arr[idxMax], arr[i]
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	results := make([]string, n)

	for daerah := 0; daerah < n; daerah++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)

		m, _ := strconv.Atoi(parts[0])

		var ganjil, genap []int

		for i := 0; i < m; i++ {
			num, _ := strconv.Atoi(parts[i+1])
			if num%2 == 1 { // Ganjil
				ganjil = append(ganjil, num)
			} else { // Genap
				genap = append(genap, num)
			}
		}

		selectionSortAscending(ganjil)

		selectionSortDescending(genap)

		var result strings.Builder
		for i, val := range ganjil {
			if i > 0 {
				result.WriteString(" ")
			}
			result.WriteString(strconv.Itoa(val))
		}
		for i, val := range genap {
			if len(ganjil) > 0 && i == 0 {
				result.WriteString(" ")
			} else if i > 0 {
				result.WriteString(" ")
			}
			result.WriteString(strconv.Itoa(val))
		}
		results[daerah] = result.String()
	}

	for _, res := range results {
		fmt.Println(res)
	}
}