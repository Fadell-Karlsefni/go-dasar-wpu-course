package main

import "fmt"

func main() {
	arr := [6]int{10, 20, 30, 40, 50, 60}

	slice1 := arr[:]
	fmt.Println("Menampilkan Semua Nilai : " , slice1) // mengambil semua nilai 

	slice2 := arr[1:3]
	fmt.Println("Mengambil Beberapa nilai : " , slice2) // mengambil beberapa nilai

	slice3 := arr[2:]
	fmt.Println("Mengambil Nilai sampai yang terakhir : " , slice3) // mengambil nilai sampai terakhir
}