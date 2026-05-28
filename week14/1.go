package main

import "fmt"

func main() {
    a := []int{64, 25, 12, 22, 11}
    n := len(a)

    i := 1
    for i <= n-1 {
        idx_min := i - 1
        j := i

        for j < n {
            if a[idx_min] > a[j] {
                idx_min = j
            }
            j = j + 1
        }

        // swap
        t := a[idx_min]
        a[idx_min] = a[i-1]
        a[i-1] = t

        i = i + 1
    }

    fmt.Println("Hasil sorting:", a)
}