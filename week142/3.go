package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func median(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n%2 == 1 {
		return arr[n/2]
	}
	return (arr[n/2-1] + arr[n/2]) / 2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var data []int

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())

		if num == -5313 {
			break
		}

		if num == 0 {
			insertionSort(data)
			fmt.Println(median(data))
		} else {
			data = append(data, num)
		}
	}
}