# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Aulia Akbar] - [109082530019]</p>

## Unguided 

### 1. [Soal 1]
#### 1.go

```go
package main

import "fmt"

func main() {
    a := []int{64, 25, 12, 22, 11}
    n := len(a)

    i := 1
    for i <= n-1 {
        idx_min := i - 1
        j := i

        for j < n {
            if a[idx_min] > a[j] {
                idx_min = j
            }
            j = j + 1
        }

        // swap
        t := a[idx_min]
        a[idx_min] = a[i-1]
        a[i-1] = t

        i = i + 1
    }

    fmt.Println("Hasil sorting:", a)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week14/output/soal1.png)
# penjelasan
Program tersebut menjelaskan tentang cara mencari nilai ekstrem pada sekumpulan data. saya menaruh data tersebut disebuah array dan nantinya saat running langsung keluar outputnya yaitu angka yang terdapat didalam array itu sendiri.

### 2. [Soal 2]
#### 2.go

```go
package main
import "fmt"

type arrInt [4321]int

func selectionSort1(T *arrInt, n int) {
	var t, i, j, idx_min int

	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

type arrMhs [2023]mahasiswa

func selectionSort2(T *arrMhs, n int) {
	var i, j, idx_min int
	var t mahasiswa

	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if T[idx_min].ipk > T[j].ipk {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func tampilkanArrayInt(T arrInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", T[i])
	}
	fmt.Println()
}

func tampilkanMahasiswa(T arrMhs, n int) {
	fmt.Printf("%-15s %-12s %-10s %-15s %s\n", "Nama", "NIM", "Kelas", "Jurusan", "IPK")
	fmt.Println("----------------------------------------------------------------")
	for i := 0; i < n; i++ {
		fmt.Printf("%-15s %-12s %-10s %-15s %.2f\n",
			T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}
}

func main() {
	fmt.Println("=== SELECTION SORT UNTUK ARRAY INTEGER ===")
	var dataInt arrInt
	nInt := 7
	dataInt[0] = 64
	dataInt[1] = 25
	dataInt[2] = 12
	dataInt[3] = 22
	dataInt[4] = 11
	dataInt[5] = 90
	dataInt[6] = 5

	fmt.Println("Array sebelum diurutkan:")
	tampilkanArrayInt(dataInt, nInt)

	selectionSort1(&dataInt, nInt)

	fmt.Println("Array setelah diurutkan secara ascending:")
	tampilkanArrayInt(dataInt, nInt)

	fmt.Println("\n")

	fmt.Println("=== SELECTION SORT UNTUK ARRAY MAHASISWA (BERDASARKAN IPK) ===")
	var dataMhs arrMhs
	nMhs := 5

	dataMhs[0] = mahasiswa{"Andi", "220001", "A", "Informatika", 3.75}
	dataMhs[1] = mahasiswa{"Budi", "220002", "B", "Sistem Informasi", 3.50}
	dataMhs[2] = mahasiswa{"Citra", "220003", "A", "Informatika", 3.90}
	dataMhs[3] = mahasiswa{"Dewi", "220004", "C", "Manajemen", 3.20}
	dataMhs[4] = mahasiswa{"Eka", "220005", "B", "Sistem Informasi", 3.65}

	fmt.Println("Data mahasiswa sebelum diurutkan (berdasarkan IPK ascending):")
	tampilkanMahasiswa(dataMhs, nMhs)

	selectionSort2(&dataMhs, nMhs)

	fmt.Println("\nData mahasiswa setelah diurutkan berdasarkan IPK (dari terendah ke tertinggi):")
	tampilkanMahasiswa(dataMhs, nMhs)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week14/output/soal2.png)
# penjelasan
Program ini menjelaskan tentang menggunakan selection sort untuk integer dan untuk struck. disoal ini saya disuruh untuk melakukan penyortiran terhadap data mahasiswa yang belum urut dengan menggunakan selection sort.

### 3. [Soal 3]
#### 3.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week14/output/soal3.png)
# penjelasan
Program tersebut mejelaskan tentang cara mengurutkan nomor sebuah kerabat disuatu daerah menggunakan selection sort.

### 4. [Soal 4]
#### 4.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week14/output/soal4.png)
# penjelasan
Program tersebut menjelaskan tentang cara kita mengurutkan sebuah nomor rumah yang awalanya dari yang ganjil terurut membesar dahulu sampai yang genap mengecil.

### 5. [Soal 5]
#### 5.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week14/output/soal5.png)
# penjelasan
Program tersebut menjelaskan tentang mencari median dari sebuah nilai nilai acak.
