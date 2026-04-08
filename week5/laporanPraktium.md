# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Aulia Akbar] - [109082530019]</p>

## Unguided 

### 1. [Soal 1]
#### 1.go

```go
ppackage main
package main
import "fmt"
func fibonacci(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
func main() {
    fmt.Println("n\t| S_n")
    fmt.Println("----------------")
    for i := 0; i <= 10; i++ {
        fmt.Printf("%d\t| %d\n", i, fibonacci(i))
    }
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week5/output/soal1.png)
# penjelasan
kodingan tersebut menjelaskan tentang deret Fibonacci yang menggunakan rumus Sn = Sn−1 + Sn−2.
### 2. [Soal 2]
#### 2.go

```go
package main

import "fmt"

func cetakBintang(n int) {
	if n == 0 {
		return
	}
	cetakBintang(n - 1)

	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	cetakBintang(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week5/output/soal 3.png)
# penjelasan
kodingan tersebut menjelaskan tentang cara kita membuat segitiga sama kaki yang awal bentuknya dari pola bintang, disini saya menggunakan func sebagai tempat menaruh variabelnya yang berisiskan for loop , dan if else.


### 3. [Soal 3]
#### 3.go

```go
package main
import "fmt"
func faktor(n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktor(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	fmt.Print("Faktor dari ", n, " adalah: ")
	faktor(n, 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week5/output/soal 3.png)
# penjelasan
kodingan tersebut adalah cara kita mencari nilai maksimum dari sederet angka/ bisa disebut angka yang paling sering muncul(modus).