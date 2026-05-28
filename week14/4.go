package main
import "fmt"

func selectionSortAsc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idx_min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idx_min] {
				idx_min = j
			}
		}
		arr[i], arr[idx_min] = arr[idx_min], arr[i]
	}
}

func selectionSortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idx_max := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[idx_max] {
				idx_max = j
			}
		}
		arr[i], arr[idx_max] = arr[idx_max], arr[i]
	}
}

func main() {
	var numDaerah int
	fmt.Scan(&numDaerah)

	for daerah := 1; daerah <= numDaerah; daerah++ {
		var n int
		fmt.Scan(&n)

		var ganjil, genap []int

		for i := 0; i < n; i++ {
			var num int
			fmt.Scan(&num)

			if num%2 == 1 {
				ganjil = append(ganjil, num)
			} else { 
				genap = append(genap, num)
			}
		}

		selectionSortAsc(ganjil)

		selectionSortDesc(genap)

		for i := 0; i < len(ganjil); i++ {
			fmt.Printf("%d ", ganjil[i])
		}
		for i := 0; i < len(genap); i++ {
			fmt.Printf("%d ", genap[i])
		}
		fmt.Println() // Pindah baris untuk daerah berikutnya
	}
}