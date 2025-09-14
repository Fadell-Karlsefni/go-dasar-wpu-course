package main

import "fmt"

func tambah(a int, b int) (hasil int) {
	hasil = a + b
	return
}

func main() {

	fmt.Println("hasil nya adalah : ", tambah(10, 5)) // hasil nya 15

}