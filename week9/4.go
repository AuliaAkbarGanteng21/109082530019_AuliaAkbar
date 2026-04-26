package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
const NMAX int = 127
type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Teks : ")
	
	if scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")
		line = strings.TrimSuffix(line, ".")
		
		*n = 0
		for i, ch := range line {
			if i < NMAX {
				t[*n] = ch
				*n++
			}
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		temp := t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrome(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	cetakArray(tab, m)

	balikanArray(&tab, m)

	cetakArray(tab, m)

	fmt.Println("\n=== Pemeriksaan Palindrome ===")
	
	var tab2 tabel
	var n int
	
	isiArray(&tab2, &n)
	
	fmt.Print("Teks : ")
	cetakArray(tab2, n)
	
	isPalindrom := palindrome(tab2, n)
	fmt.Printf("Palindrome : %t\n", isPalindrom)
}