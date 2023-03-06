package main

import "fmt"

func HitungKembalian(totalbeli, bayar int) int {
	fmt.Print("Total Beli : ")
	fmt.Scanln(&totalbeli)

	fmt.Print("Bayar : ")
	fmt.Scanln(&bayar)

	kembalian := bayar - totalbeli
	return kembalian
}

func HitungPecahan(kembalian *int, pec20k, pec10k, pec5k, pec2k, pec1k *int) {
	var totalbeli, bayar int
	*kembalian = HitungKembalian(totalbeli, bayar)
	fmt.Println("Kembalian :",*kembalian)

	for *kembalian != 0 {
		if *kembalian >= 20000 {
			*pec20k = *kembalian / 20000
			*kembalian = *kembalian - 20000
		} else if *kembalian >= 10000 && *kembalian < 20000 {
			*pec10k = *kembalian / 10000
			*kembalian = *kembalian - 10000
		} else if *kembalian >= 5000 && *kembalian < 10000 {
			*pec5k = *kembalian / 5000
			*kembalian = *kembalian - 5000
		} else if *kembalian >= 2000 && *kembalian < 5000 {
			*pec2k = *kembalian / 2000
			*kembalian = *kembalian - 2000
		} else if *kembalian >= 1000 && *kembalian < 5000 {
			*pec1k = *kembalian / 1000
			*kembalian = *kembalian - 1000
		}
	}

	fmt.Println("Pecahan Rp.20.000,- sebanyak", *pec20k, "buah")
	fmt.Println("Pecahan Rp.10.000,- sebanyak", *pec10k, "buah")
	fmt.Println("Pecahan Rp.5.000,- sebanyak", *pec5k, "buah")
	fmt.Println("Pecahan Rp.2.000,- sebanyak", *pec2k, "buah")
	fmt.Println("Pecahan Rp.1.000,- sebanyak", *pec1k, "buah")
}

func main() {
	var kembalian int
	var pec20k, pec10k, pec5k, pec2k, pec1k int
	HitungPecahan(&kembalian, &pec20k, &pec10k, &pec5k, &pec2k, &pec1k)
}
