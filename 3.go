package main

import "fmt"

const NMAX int = 100

type buku struct{
	id,judul,penerbit string
	jumlah,tahun,rating int
}

type tabBuku[NMAX]buku

func  inputBuku(T *tabBuku, N *int){
	var id,judul,penerbit string
	var jumlah,tahun,rating int

	for id!="STOP"{
		fmt.Scan(&id)
		if id=="STOP" {break}
		fmt.Scan(&judul)
		fmt.Scan(&penerbit)
		fmt.Scan(&jumlah)
		fmt.Scan(&tahun)
		fmt.Scan(&rating)
		(*T)[*N].id=id
		(*T)[*N].judul=judul
		(*T)[*N].penerbit=penerbit
		(*T)[*N].jumlah=jumlah
		(*T)[*N].tahun=tahun
		(*T)[*N].rating=rating
		*N++
	} 
}	

func print(T *tabBuku, N *int){
	for i := 0; i < *N; i++ {
		fmt.Print((*T)[i].id, " ")
		fmt.Print((*T)[i].judul, " ")
		fmt.Print((*T)[i].penerbit, " ")
		fmt.Print((*T)[i].jumlah, " ")
		fmt.Print((*T)[i].tahun, " ")
		fmt.Println((*T)[i].rating)
	}
}

func Terfavorit(T *tabBuku, N int){
	for i := 0; i < N; i++ {
		for j:=0;j<N;j++{
			if (*T)[i].rating > (*T)[j].rating{
				(*T)[i],(*T)[j] =(*T)[j],(*T)[i]  
			}
		}
	}
	fmt.Println("Terfavorit")
	fmt.Println("ID :",(*T)[0].id)
	fmt.Println("Judul :",(*T)[0].judul)
	fmt.Println("Penerbit :",(*T)[0].penerbit)
	fmt.Println("Jumlah :",(*T)[0].jumlah)
	fmt.Println("Tahun :",(*T)[0].tahun)
	fmt.Println("Rating :",(*T)[0].rating)
}

func UrutBuku(T *tabBuku, N int){
	for i := 0; i < N; i++ {
		for j:=0;j<N;j++{
			if (*T)[i].tahun > (*T)[j].tahun{
				(*T)[i],(*T)[j] =(*T)[j],(*T)[i]  
			}
		}
	}
	print(T, &N)
}

func LimaTerbaru(T *tabBuku, N int){
	var limit int
	for i := 0; i < N; i++ {
		for j:=0;j<N;j++{
			if (*T)[i].tahun > (*T)[j].tahun{
				(*T)[i],(*T)[j] =(*T)[j],(*T)[i]  
			}
		}
	}
	
	limit = 5
	if N < limit {
		fmt.Println("Buku kurang dari 5.","Namun akan ditampilkan", N, "buku terbaru")
		for k:=0;k<N;k++{
			fmt.Println((*T)[k].judul)
		}
	} else {
		for k:=0;k<limit;k++{
			fmt.Println((*T)[k].judul)
		}
	}
} 

func main(){
	var N int
	var T tabBuku
	inputBuku(&T, &N)

	Terfavorit(&T,N)
	fmt.Println()

	UrutBuku(&T, N)
	fmt.Println()
	
	LimaTerbaru(&T, N)
}