package main
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