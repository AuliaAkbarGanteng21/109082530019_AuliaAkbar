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