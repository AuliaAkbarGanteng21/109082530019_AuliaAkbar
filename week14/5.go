package main
import "fmt"

func insertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i - 1

		for j >= 0 && data[j] > temp {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

func median(data []int) int {
	insertionSort(data)

	n := len(data)

	if n%2 == 1 {
		return data[n/2]
	}
	return (data[n/2-1] + data[n/2]) / 2
}

func main() {
	var x int
	var data []int

	for {
		fmt.Scan(&x)

		if x == -5313 {
			break
		}

		if x == 0 {
			fmt.Println(median(data))
		} else {
			data = append(data, x)
		}
	}
}