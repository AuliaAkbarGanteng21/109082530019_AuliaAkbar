package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	if n == 0 {
		return 0.0
	}

	var total float64 = 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

type HasilPosyandu struct {
	Minimum  float64
	Maximum  float64
	Rerata   float64
	JumlahData int
}

func main() {
	var data arrBalita
	var n int
	var min, max float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	if n > 100 {
		fmt.Println("Maksimal data adalah 100")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	hitungMinMax(data, n, &min, &max)

	rata := rerata(data, n)

	hasil := HasilPosyandu{
		Minimum:    min,
		Maximum:    max,
		Rerata:     rata,
		JumlahData: n,
	}

	fmt.Printf("Berat balita minimum: %.2f kg\n", hasil.Minimum)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", hasil.Maximum)
	fmt.Printf("Rerata berat balita: %.2f kg\n", hasil.Rerata)
}