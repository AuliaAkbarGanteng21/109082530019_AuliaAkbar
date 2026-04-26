package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N)

	// Membuat array dengan kapasitas N
	arr := make([]int, N)

	// Mengisi array
	fmt.Println("Masukkan", N, "bilangan bulat:")
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	// Menu pilihan
	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Tampilkan seluruh isi array")
		fmt.Println("2. Tampilkan elemen indeks ganjil")
		fmt.Println("3. Tampilkan elemen indeks genap")
		fmt.Println("4. Tampilkan elemen indeks kelipatan x")
		fmt.Println("5. Hapus elemen pada indeks tertentu")
		fmt.Println("6. Tampilkan rata-rata")
		fmt.Println("7. Tampilkan standar deviasi")
		fmt.Println("8. Tampilkan frekuensi suatu bilangan")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu (1-9): ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilkanSemua(arr)
		case 2:
			tampilkanIndeksGanjil(arr)
		case 3:
			tampilkanIndeksGenap(arr)
		case 4:
			fmt.Print("Masukkan nilai x: ")
			var x int
			fmt.Scan(&x)
			tampilkanKelipatan(arr, x)
		case 5:
			fmt.Print("Masukkan indeks yang akan dihapus: ")
			var idx int
			fmt.Scan(&idx)
			arr = hapusElemen(arr, idx)
			fmt.Println("Array setelah penghapusan:")
			tampilkanSemua(arr)
		case 6:
			rata := hitungRataRata(arr)
			fmt.Printf("Rata-rata: %.2f\n", rata)
		case 7:
			stdDev := hitungStandarDeviasi(arr)
			fmt.Printf("Standar deviasi: %.2f\n", stdDev)
		case 8:
			fmt.Print("Masukkan bilangan yang dicari: ")
			var target int
			fmt.Scan(&target)
			freq := hitungFrekuensi(arr, target)
			fmt.Printf("Frekuensi bilangan %d: %d kali\n", target, freq)
		case 9:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

// a. Menampilkan keseluruhan isi array
func tampilkanSemua(arr []int) {
	fmt.Print("Isi array: ")
	for i, val := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()
}

// b. Menampilkan elemen dengan indeks ganjil
func tampilkanIndeksGanjil(arr []int) {
	fmt.Print("Elemen indeks ganjil: ")
	pertama := true
	for i := 1; i < len(arr); i += 2 {
		if !pertama {
			fmt.Print(" ")
		}
		fmt.Print(arr[i])
		pertama = false
	}
	if pertama {
		fmt.Print("(tidak ada)")
	}
	fmt.Println()
}

// c. Menampilkan elemen dengan indeks genap (indeks 0 = genap)
func tampilkanIndeksGenap(arr []int) {
	fmt.Print("Elemen indeks genap: ")
	pertama := true
	for i := 0; i < len(arr); i += 2 {
		if !pertama {
			fmt.Print(" ")
		}
		fmt.Print(arr[i])
		pertama = false
	}
	fmt.Println()
}

// d. Menampilkan elemen dengan indeks kelipatan x
func tampilkanKelipatan(arr []int, x int) {
	if x <= 0 {
		fmt.Println("x harus lebih besar dari 0")
		return
	}
	fmt.Printf("Elemen indeks kelipatan %d: ", x)
	pertama := true
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			if !pertama {
				fmt.Print(" ")
			}
			fmt.Print(arr[i])
			pertama = false
		}
	}
	if pertama {
		fmt.Print("(tidak ada)")
	}
	fmt.Println()
}

// e. Menghapus elemen pada indeks tertentu
func hapusElemen(arr []int, idx int) []int {
	if idx < 0 || idx >= len(arr) {
		fmt.Println("Indeks tidak valid!")
		return arr
	}
	// Menghapus elemen dengan slicing
	return append(arr[:idx], arr[idx+1:]...)
}

// f. Menghitung rata-rata
func hitungRataRata(arr []int) float64 {
	if len(arr) == 0 {
		return 0
	}
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return float64(sum) / float64(len(arr))
}

// g. Menghitung standar deviasi
func hitungStandarDeviasi(arr []int) float64 {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return 0
	}

	rata := hitungRataRata(arr)
	var sumSquares float64

	for _, val := range arr {
		diff := float64(val) - rata
		sumSquares += diff * diff
	}

	variance := sumSquares / float64(len(arr))
	return math.Sqrt(variance)
}

// h. Menghitung frekuensi suatu bilangan
func hitungFrekuensi(arr []int, target int) int {
	count := 0
	for _, val := range arr {
		if val == target {
			count++
		}
	}
	return count
}