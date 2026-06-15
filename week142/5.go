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
			high = mid - 1 // karena descending, rating lebih besar di kiri
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