# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Aulia Akbar] - [109082530019]</p>

## Unguided 

### 1. [Soal 1]
#### 1.go

```go
package main
import "fmt"
func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}


func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int
	
	fmt.Scan(&a, &b, &c, &d)
	
	p1 := permutasi(a, c)
	c1 := kombinasi(a, c)
	fmt.Printf("%d %d\n", p1, c1)
	
	p2 := permutasi(b, d)
	c2 := kombinasi(b, d)
	fmt.Printf("%d %d\n", p2, c2)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week4/output/soal1.png)
# penjelasan
Program tersebut menjelaskan tentang cara menghitung soal matematika yang diimplementasikan oleh seorang mahasiswa yang berisikan tentang menghitung 4 bilangan yang berisikan bilangan a, b, c, d yang dimana program tersebut bersyarat a > c dan b > d.
### 2. [Soal 2]
#### 2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week4/output/soal2.png)
# penjelasan
Program ini menjelaskan tentang cara menentukan pemenang lomba programming yang dimana peserta peserta diberikan 8 soal , dan dari situ kalau peserta yang bisa menyelesaikan soal tersebut <301 dianggap selesai , dan kalau peserta yang > 301 dianggap gagal. dan dihitung dari segi pememenangnya kalau peserta tersebut menyelesaikan soal paling banyak / yang bisa mengerjakan semuanya, dan bila sama dihitungnya dari waktu total penyelesaiannya yang paling kecil.