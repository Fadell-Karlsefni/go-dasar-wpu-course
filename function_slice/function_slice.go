package main

import "fmt"

func main() {
	slice := make([]int, 4, 5)
	fmt.Println(slice)

	fmt.Println("Panjang slice nya : ", len(slice)) // menggunakan len untuk mengetahui panjang slice
	fmt.Println("Capasitas slice nya : ", cap(slice)) // menggunakan cap untuk cek capasitas slice

	slice = append(slice, 20,30,40)
	fmt.Println("append menambahkan nilai : ",slice) // menambahkan nilai ke dalam slice

	slice2 := make([]int, 7)
	slice3 := copy(slice2,slice)
	fmt.Println("copy : ", slice2) // copy array sebelum nya
	fmt.Println("Jumlah elemen yang terasalin : ", slice3)

	angka := []int {1,2,3,4,5,6}
	slice4 := angka[1:4]
	fmt.Println("ini adalah membuat slice di dalam array : ", slice4 )
}