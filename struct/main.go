package main

import "fmt"

type persegiPanjang struct {
	panjang, lebar float64
}

// method menghitung luas
func (p persegiPanjang) luas() float64 {
	return p.panjang * p.lebar
}

func main() {
	pp := persegiPanjang{panjang: 10, lebar: 5}
	pp2 := persegiPanjang{panjang: 15, lebar: 5}
	fmt.Println("<==Menghitung persegi yang pertama==>")
	fmt.Println("Panjang : ", pp.panjang)
	fmt.Println("Lebar : ", pp.lebar)
	fmt.Println("luas : ", pp.luas())
	fmt.Println("<==Menghitung persegi yang Kedua==>")
	fmt.Println("Panjang : ", pp2.panjang)
	fmt.Println("Lebar : ", pp2.lebar)
	fmt.Println("luas : ", pp2.luas())
}