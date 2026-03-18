# <h1 align="center">Laporan Praktikum Modul 3 - ... </h1>
<p align="center">[Aulia Akbar] - [109082530019]</p>

## Unguided 

### 1. [Soal 1]
#### 1.go

```go
package main
 import "fmt"
 func factorial(n int) int {
 	result := 1
 	for i := 2; i <= n; i++ {
 		result *= i
 	}
 	return result
 }

 func permutation(n, r int) int {
 	return factorial(n) / factorial(n-r)
 }
 func combination(n, r int) int {
 	return factorial(n) / (factorial(r) * factorial(n-r))
 }
 func main() {
 	var a, b, c, d int
 	fmt.Scan(&a, &b, &c, &d)
 	fmt.Printf("%d %d\n", permutation(a, c), combination(a, c))
 	fmt.Printf("%d %d\n", permutation(b, d), combination(b, d))
 }
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week3/output/soal1.png)
# penjelasan
kodingan tersebut menjelaskan menghitung dan menampikan hasil permutasian serta kombinasi dari dua pasangan bilangan bulat yang dimasukan pengguna.
### 2. [Soal 2]
#### 2.go

```go
package main
import "fmt"
func f(x int) int {
	return x * x
}
func g(x int) int {
	return x - 2
}
func h(x int) int {
	return x + 1
}
func fogoh(x int) int {
	return f(g(h(x)))
}
func gohof(x int) int {
	return g(h(f(x)))
}
func hofog(x int) int {
	return h(f(g(x)))
}
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(fogoh(a))
	fmt.Println(gohof(b))
	fmt.Println(hofog(c))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week3/output/soal 3.png)
# penjelasan
kodingan tersebut menjelaskan cara mengimplementasikan konsep komposisi fungsi dalam ilmu matematika dan menghitung hasilnya secara otomatis untuk nilai input yang diberikan.


### 3. [Soal 3]
#### 3.go

```go
package main
import (
	"fmt"
	"math"
)
func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}
func didalam(cx, cy, r, x, y float64) bool {
	return jarak(x, y, cx, cy) <= r
}
func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	dalam1 := didalam(cx1, cy1, r1, x, y)
	dalam2 := didalam(cx2, cy2, r2, x, y)

	switch {
	case dalam1 && dalam2:
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case dalam1:
		fmt.Println("Titik di dalam lingkaran 1")
	case dalam2:
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/AuliaAkbarGanteng21/109082530019_AuliaAkbar/blob/main/week3/output/soal 3.png)
# penjelasan
dikodingan tersebut menjelaskan cara menentukan posisi sebuah titik sembarang terhadap dua lingkaran yang telah didefinisikan. program ini menerapkan konsep geometri dasar terkait dengan lingkaran dan jarak antar titik. 