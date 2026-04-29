# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Aulia Akbar] - [109082530019]</p>

## Unguided 

### 1. [Soal 1]
#### 1.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	if n <= 0 || n > 1000 {
		fmt.Println("Jumlah data harus antara 1 - 1000")
		return
	}

	var berat [1000]float64

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]

	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Println("Berat terkecil:", min)
	fmt.Println("Berat terbesar:", max)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week10/output/soal1.png)
# penjelasan
Program tersebut menjelaskan tentang cara kita menghitung berat dari sebuah kelinci dengan menggunakan arrray dengan kapasitas 1000, bilangan bulat yang digunakan untuk menimbang berat anak kelinci, bilangan rill digunakan untuk berat anak kelinci yang akan dijual.

### 2. [Soal 2]
#### 2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week10/output/soal2.png)
# penjelasan
Program tersebut menjelaskan tentang cara kita menghitung berat ikan menggunakan Array dengan kapasitas berat ikan 1000 untuk menampung data berat ikan yang akan dijual.menggunakan 2 bilangan bulat yaitu x dan y. x menyatakan bilangan berat ikan yang akan dijual dan y adalah bilangan berat ikan yang akan ditaro wadah.



### 3. [Soal 3]
#### 3.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week10/output/soal3.png)
# penjelasan
1. type arrBalita [100]float64 — Mendefinisikan tipe array dengan kapasitas 100 (sesuai spesifikasi).
2. Prosedur hitungMinMax — Menerima array, jumlah data (n), dan dua parameter pointer (bMin, bMax). Prosedur ini mencari nilai terkecil dan terbesar, kemudian menyimpannya ke alamat memori yang ditunjuk oleh pointer.
3. Fungsi rerata — Menerima array dan jumlah data, mengembalikan nilai rata-rata bertipe float64.
4. struct HasilPosyandu — Struct yang digunakan untuk menyimpan hasil perhitungan (minimum, maksimum, rerata, dan jumlah data). Menunjukkan pemahaman tentang penggunaan struct.
5. Validasi — Memastikan jumlah data n tidak melebihi kapasitas array (100)