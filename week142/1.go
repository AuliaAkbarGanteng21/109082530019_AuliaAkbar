package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// selectionSort mengurutkan slice secara ascending menggunakan algoritma selection sort
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		// Tukar elemen
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	// Baca n (jumlah daerah)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	
	results := make([]string, n)
	
	for daerah := 0; daerah < n; daerah++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)
		
		// Bagian pertama adalah m (jumlah kerabat di daerah ini)
		m, _ := strconv.Atoi(parts[0])
		
		// Baca m bilangan berikutnya
		rumah := make([]int, m)
		for i := 0; i < m; i++ {
			rumah[i], _ = strconv.Atoi(parts[i+1])
		}
		
		// Urutkan menggunakan selection sort
		selectionSort(rumah)
		
		// Simpan hasil dalam bentuk string
		var result strings.Builder
		for i, val := range rumah {
			if i > 0 {
				result.WriteString(" ")
			}
			result.WriteString(strconv.Itoa(val))
		}
		results[daerah] = result.String()
	}
	
	// Cetak semua hasil
	for _, res := range results {
		fmt.Println(res)
	}
}