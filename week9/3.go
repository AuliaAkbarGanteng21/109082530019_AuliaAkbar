package main

import "fmt"

func main() {
	var klubA, klubB string

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	pemenang := []string{}

	pertandinganKe := 1

	for {
		var skorA, skorB int

		// Memasukkan skor pertandingan
		fmt.Printf("Pertandingan %d : ", pertandinganKe)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
			fmt.Printf("Hasil %d : %s\n", pertandinganKe, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
			fmt.Printf("Hasil %d : %s\n", pertandinganKe, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
			fmt.Printf("Hasil %d : Draw\n", pertandinganKe)
		}

		pertandinganKe++
	}

	fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
	for i, p := range pemenang {
		fmt.Printf("%d. %s\n", i+1, p)
	}
}