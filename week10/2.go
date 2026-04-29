package main

import "fmt"

type Wadah struct {
	total float64
}

func hitungWadah(ikan [1000]float64, x, y int, wadah *[1000]Wadah) int {
	jumlahWadah := (x + y - 1) / y
	idx := 0

	for i := 0; i < jumlahWadah; i++ {
		for j := 0; j < y && idx < x; j++ {
			wadah[i].total += ikan[idx]
			idx++
		}
	}
	return jumlahWadah
}

func rataWadah(wadah [1000]Wadah, n int) float64 {
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += wadah[i].total
	}
	return sum / float64(n)
}

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	var ikan [1000]float64
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	var wadah [1000]Wadah

	jumlahWadah := hitungWadah(ikan, x, y, &wadah)

	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("%.2f ", wadah[i].total)
	}

	fmt.Printf("\n%.2f\n", rataWadah(wadah, jumlahWadah))
}