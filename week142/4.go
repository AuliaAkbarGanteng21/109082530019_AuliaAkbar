package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	input = strings.ReplaceAll(input, " ", "")

	parts := strings.Split(input, ",")

	var data []int
	for _, p := range parts {
		num, err := strconv.Atoi(p)
		if err != nil {
			continue
		}
		if num < 0 {
			break
		}
		data = append(data, num)
	}

	insertionSort(data)

	for i, val := range data {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(val)
	}
	fmt.Println()

	if len(data) < 2 {
		fmt.Println("Data berjarak tidak tetap")
		return
	}

	jarak := data[1] - data[0]
	berjarakTetap := true
	for i := 2; i < len(data); i++ {
		if data[i]-data[i-1] != jarak {
			berjarakTetap = false
			break
		}
	}

	if berjarakTetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}