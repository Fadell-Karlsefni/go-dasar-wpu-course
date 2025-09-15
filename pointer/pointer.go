package main

import "fmt"

func ubahNama(nama *string) {
	*nama = "budi"
}

func main() {
	nama := "andi"
	ubahNama(&nama)
	fmt.Println(nama)
}