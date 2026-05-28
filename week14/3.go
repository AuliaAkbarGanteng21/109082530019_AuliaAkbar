package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		a := make([]int, m)

		for j := 0; j < m; j++ {
			fmt.Scan(&a[j])
		}

		for j := 0; j < m-1; j++ {
			min := j

			for k := j + 1; k < m; k++ {
				if a[k] < a[min] {
					min = k
				}
			}

			a[j], a[min] = a[min], a[j]
		}

		for j := 0; j < m; j++ {
			fmt.Print(a[j], " ")
		}
		fmt.Println()
	}
}