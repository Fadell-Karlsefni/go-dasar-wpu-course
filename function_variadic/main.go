package main

import "fmt"

func jumlahAngka(angka ...int) int {
	hasil := 0
	for _, v := range angka {
		hasil += v
	}
	return hasil
}

func main() {

	fmt.Println(jumlahAngka(2,3,4))
	fmt.Println(jumlahAngka(10,30,45))
	fmt.Println(jumlahAngka())

}