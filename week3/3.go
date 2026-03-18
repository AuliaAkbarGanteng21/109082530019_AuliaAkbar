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