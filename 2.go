package main

import "fmt"

type poin struct {
	nilai float64
}

const NMAX = 999

type arrNilai [NMAX]poin

func inputNilai(T *arrNilai, n *int) {
	var input float64
	fmt.Scan(&input)
	for input != -1 {
		(*T)[*n].nilai = input
		*n++
		fmt.Scan(&input)
	}
}

func rataRata(T *arrNilai, n *int) float64 {
	var total, nilai float64
	for i := 0; i < *n; i++ {
		total += ((*T)[i].nilai)
	}
	nilai = total / float64(*n)
	return nilai
}

func minMax(T *arrNilai, n *int, minValue,maxValue *float64) {
	for i := 0; i < *n; i++ {
		for j := i; j < *n; j++ {
			if (*T)[j].nilai > (*T)[i].nilai {
				(*T)[j], (*T)[i] = (*T)[i], (*T)[j]
			}
		}
	}
	*minValue = (*T)[*n-1].nilai
	*maxValue = (*T)[0].nilai

}

func main() {
	var T arrNilai
	var jumlah int
	var average, minValue, maxValue float64

	inputNilai(&T, &jumlah)
	average = (rataRata(&T, &jumlah))
	minMax(&T,&jumlah,&minValue,&maxValue)
	
	fmt.Println("Banyak nilai :", jumlah)
	fmt.Println("Nilai rata-rata :", average)
	fmt.Println("Nilai maximum :", maxValue)
	fmt.Println("Nilai minimum :", minValue)
}
