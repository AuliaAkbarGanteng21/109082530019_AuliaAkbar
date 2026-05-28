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