# <h1 align="center">Laporan Praktikum Modul 14 - ... </h1>
<p align="center">[Aulia Akbar] - [109082530019]</p>

## Unguided 

### 1. [Soal 1]
#### 1.go

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// selectionSort mengurutkan slice secara ascending menggunakan algoritma selection sort
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		// Tukar elemen
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	// Baca n (jumlah daerah)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	
	results := make([]string, n)
	
	for daerah := 0; daerah < n; daerah++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)
		
		// Bagian pertama adalah m (jumlah kerabat di daerah ini)
		m, _ := strconv.Atoi(parts[0])
		
		// Baca m bilangan berikutnya
		rumah := make([]int, m)
		for i := 0; i < m; i++ {
			rumah[i], _ = strconv.Atoi(parts[i+1])
		}
		
		// Urutkan menggunakan selection sort
		selectionSort(rumah)
		
		// Simpan hasil dalam bentuk string
		var result strings.Builder
		for i, val := range rumah {
			if i > 0 {
				result.WriteString(" ")
			}
			result.WriteString(strconv.Itoa(val))
		}
		results[daerah] = result.String()
	}
	
	// Cetak semua hasil
	for _, res := range results {
		fmt.Println(res)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week142/output/soal1.png)
# penjelasan
Buat program yang membaca bilangan bulat (dipisah koma), hanya menyimpan bilangan ≥ 0, berhenti jika bertemu bilangan negatif, lalu mengurutkan data dengan Insertion Sort, menampilkan data terurut (format dipisah koma), lalu mengecek apakah selisih antar data yang berurutan sama semua. Jika iya, cetak "Data berjarak x" (x = selisihnya); jika tidak, cetak "Data berjarak tidak tetap".
1. Hanya simpan bilangan non-negatif.
2. Urutkan dengan Insertion Sort (wajib).
3. Cek apakah selisih antar elemen berurutan konstan.
4. Keluaran sesuai format contoh.

### 2. [Soal 2]
#### 2.go

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func selectionSortAscending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func selectionSortDescending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[idxMax] {
				idxMax = j
			}
		}
		arr[i], arr[idxMax] = arr[idxMax], arr[i]
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	results := make([]string, n)

	for daerah := 0; daerah < n; daerah++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)

		m, _ := strconv.Atoi(parts[0])

		var ganjil, genap []int

		for i := 0; i < m; i++ {
			num, _ := strconv.Atoi(parts[i+1])
			if num%2 == 1 { // Ganjil
				ganjil = append(ganjil, num)
			} else { // Genap
				genap = append(genap, num)
			}
		}

		selectionSortAscending(ganjil)

		selectionSortDescending(genap)

		var result strings.Builder
		for i, val := range ganjil {
			if i > 0 {
				result.WriteString(" ")
			}
			result.WriteString(strconv.Itoa(val))
		}
		for i, val := range genap {
			if len(ganjil) > 0 && i == 0 {
				result.WriteString(" ")
			} else if i > 0 {
				result.WriteString(" ")
			}
			result.WriteString(strconv.Itoa(val))
		}
		results[daerah] = result.String()
	}

	for _, res := range results {
		fmt.Println(res)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week142/output/soal2.png)
# penjelasan
Program "kerabatdekat"
Memisahkan nomor rumah menjadi dua kelompok:
Ganjil → diurutkan membesar (ascending)
Genap → diurutkan mengecil (descending)
Algoritma yang digunakan: Selection Sort (wajib)
Format masukan:
Baris pertama = jumlah daerah (n)
Setiap daerah: m (jumlah rumah) diikuti m nomor rumah
Format keluaran:
Setiap daerah satu baris, cetak semua ganjil terurut membesar dulu, lalu semua genap terurut mengecil
Pisahkan, urutkan dengan Selection Sort (ascending untuk ganjil, descending untuk genap), lalu cetak berurutan.



### 3. [Soal 3]
#### 3.go

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func median(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n%2 == 1 {
		return arr[n/2]
	}
	return (arr[n/2-1] + arr[n/2]) / 2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var data []int

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())

		if num == -5313 {
			break
		}

		if num == 0 {
			insertionSort(data)
			fmt.Println(median(data))
		} else {
			data = append(data, num)
		}
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week142/output/soal3.png)
# penjelasan
Program "median"
Membaca bilangan bulat secara berurutan, dengan ketentuan:

Bilangan 0 → tanda untuk menghitung dan mencetak median dari semua data yang telah tersimpan sejauh ini (0 tidak ikut disimpan).
Bilangan -5313 → tanda akhir program (tidak diproses).
Bilangan positif lainnya → disimpan ke dalam array.
Aturan median:
Data diurutkan terlebih dahulu (gunakan Insertion Sort setiap kali akan menghitung median).
Jika jumlah data ganjil → median adalah nilai tengah.
Jika jumlah data genap → median adalah rata-rata dua nilai tengah, dibulatkan ke bawah.
Keluaran: Setiap kali menemukan 0, cetak median dalam satu baris.

### 4. [Soal 4]
#### 4.go

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	input = strings.ReplaceAll(input, " ", "")

	parts := strings.Split(input, ",")

	var data []int
	for _, p := range parts {
		num, err := strconv.Atoi(p)
		if err != nil {
			continue
		}
		if num < 0 {
			break
		}
		data = append(data, num)
	}

	insertionSort(data)

	for i, val := range data {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(val)
	}
	fmt.Println()

	if len(data) < 2 {
		fmt.Println("Data berjarak tidak tetap")
		return
	}

	jarak := data[1] - data[0]
	berjarakTetap := true
	for i := 2; i < len(data); i++ {
		if data[i]-data[i-1] != jarak {
			berjarakTetap = false
			break
		}
	}

	if berjarakTetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week142/output/soal4.png)
# penjelasan
Program membaca bilangan bulat secara berurutan. Bilangan positif disimpan. Bilangan 0 menjadi tanda untuk mengurutkan semua data yang sudah tersimpan (dengan Insertion Sort), lalu menghitung dan mencetak median (ganjil = nilai tengah, genap = rata-rata dua nilai tengah dibulatkan ke bawah). Bilangan -5313 mengakhiri program. Data tidak direset setelah mencetak median.


### 5. [Soal 5]
#### 5.go

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating      int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	*n, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < *n; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Split(line, ",")
		for j := range parts {
			parts[j] = strings.TrimSpace(parts[j])
		}

		eks, _ := strconv.Atoi(parts[4])
		tahun, _ := strconv.Atoi(parts[5])
		rating, _ := strconv.Atoi(parts[6])

		pustaka[i] = Buku{
			id:        parts[0],
			judul:     parts[1],
			penulis:   parts[2],
			penerbit:  parts[3],
			eksemplar: eks,
			tahun:     tahun,
			rating:    rating,
		}
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		return
	}
	idxMax := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
	}
	b := pustaka[idxMax]
	fmt.Printf("%s, %s, %s, %d\n", b.judul, b.penulis, b.penerbit, b.tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}
	for i := 0; i < limit; i++ {
		b := pustaka[i]
		fmt.Printf("%s, %s, %s, %d, %d, %d\n",
			b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	low, high := 0, n-1
	found := -1
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			found = mid
			break
		} else if pustaka[mid].rating < r {
			high = mid - 1 
		} else {
			low = mid + 1
		}
	}

	if found == -1 {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	} else {
		b := pustaka[found]
		fmt.Printf("%s, %s, %s, %d, %d, %d\n",
			b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
	}
}

func main() {
	var pustaka DaftarBuku
	var n int

	DaftarkanBuku(&pustaka, &n)

	CetakTerfavorit(pustaka, n)

	UrutBuku(&pustaka, n)

	Cetak5Terbaru(pustaka, n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	r, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	CariBuku(pustaka, n, r)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week142/output/soal5.png)
# penjelasan
Program mengelola data buku dengan struct berisi: id, judul, penulis, penerbit, eksemplar, tahun, rating.
5 subprogram wajib:
1. DaftarkanBuku – baca N dan N data buku dari input.
2. CetakTerfavorit – cari & cetak buku rating tertinggi (judul, penulis, penerbit, tahun) tanpa sorting.
3. UrutBuku – urutkan semua buku menurun berdasarkan rating menggunakan Insertion Sort.
4. Cetak5Terbaru – cetak maksimal 5 buku rating tertinggi (setelah sorting).
5. CariBuku – cari buku berdasarkan rating dengan Binary Search. Jika tidak ditemukan, cetak "Tidak ada buku dengan rating seperti itu".
Alur main: Daftar → CetakTerfavorit → Urut → Cetak5Terbaru → baca rating cari → CariBuku
Inti kunci: Insertion Sort (wajib) + Binary Search (wajib)