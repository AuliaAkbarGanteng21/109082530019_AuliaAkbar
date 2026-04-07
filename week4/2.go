package main

import (
	"bufio"
	"fmt"
	"os"
)

func hitungSkor(waktu [8]int, soal *int, skor *int) {
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		if waktu[i] < 301 {
			*soal++
			*skor += waktu[i]
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nama, namaPemenang string
	var maxSoal, minSkor int
	maxSoal = -1
	minSkor = 999999

	for {
		scanner.Scan()
		line := scanner.Text()

		
		if line == "Selesai" {
			break
		}

		var waktu [8]int

		fmt.Sscanf(line, "%s %d %d %d %d %d %d %d %d",
			&nama,
			&waktu[0], &waktu[1], &waktu[2], &waktu[3],
			&waktu[4], &waktu[5], &waktu[6], &waktu[7],
		)

		var soal, skor int
		hitungSkor(waktu, &soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			maxSoal = soal
			minSkor = skor
			namaPemenang = nama
		}
	}

	fmt.Printf("%s %d %d\n", namaPemenang, maxSoal, minSkor)
}