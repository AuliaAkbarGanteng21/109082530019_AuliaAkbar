# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Aulia Akbar] - [109082530019]</p>

## Unguided 

### 1. [Soal 1]
#### 1.go

```go
ppackage main
import "fmt"
func main() {
	var(
		satu, dua, tiga string
		temp string
	)
	fmt.Print("Masukkan kata pertama: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukkan kata kedua: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukkan kata ketiga: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week2/output/soal1.png)
# penjelasan
kodingan tersebut menjelaskan tentang sebuah sistem yang dimana sistem tersebut berisikan cara kita menyebutkan kata pertama sampai ketiga dan nantinya outputnya menjadi kata ketiga kedua dan pertama.
### 2. [Soal 2]
#### 2.go

```go
package main
import "fmt"
func main() {
	var i, j int
	var a1, a2, a3, a4 string
	var berhasil bool

	for i = 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&a1, &a2, &a3, &a4)

		if a1 == "merah" && a2 == "kuning" && a3 == "hijau" && a4 == "ungu" {
			j++
		}

	}

	berhasil = j == 5
	fmt.Println("BERHASIL:", berhasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week2/output/soal 3.png)
# penjelasan
kodingan tersebut menjelaskan tentang mengecek sebuah percobaan yang diulang sampai lima kali dan bila kata yang disebutkan diperulangan tersebut benar semua maka akan bernilai kan true dan jika ada yang salah maka akan bernilaikan false.



### 3. [Soal 3]
#### 3.go

```go
package main
import "fmt"
func main() {
	var (
		 parsel, beratsisa, d1, biayakirim, biayajasa, total int
	)

	fmt.Print("Berat Parsel (gram): ")
	fmt.Scan(&parsel)

	d1 = parsel / 1000
	beratsisa = parsel % 1000

	if beratsisa >= 500 && parsel <= 10000 {
		biayakirim = beratsisa * 5
	} else if beratsisa < 500 && parsel <= 10000 {
		biayakirim = beratsisa * 15
	} else {
		biayakirim = beratsisa * 0
	}

	biayajasa = d1 * 10000
	total = biayajasa + biayakirim

	fmt.Printf("Detail berat: %d kg + %d gr\n", d1, beratsisa)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayajasa, biayakirim)
	fmt.Printf("Total biaya pengiriman: Rp. %d\n", total)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week2/output/soal 3.png)
# penjelasan
dikodingan tersebut menjelaskan cara kerja aplikasi yang menghitung biaya pengiriman bedasarkan sebuah parsel. yang menggunakan perulangan if dan else dan semuaanya dikalikan dengan 10000.