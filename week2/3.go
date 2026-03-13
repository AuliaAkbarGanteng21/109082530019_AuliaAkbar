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