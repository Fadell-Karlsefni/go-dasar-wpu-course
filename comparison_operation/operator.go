package main

import "fmt"

func main() {
	var a, b int

	// Input dari terminal
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scan(&b)

	fmt.Println("\n=== Hasil Perbandingan ===")
	fmt.Printf("%d == %d : %t\n", a, b, a == b)
	fmt.Printf("%d != %d : %t\n", a, b, a != b)
	fmt.Printf("%d >  %d : %t\n", a, b, a > b)
	fmt.Printf("%d <  %d : %t\n", a, b, a < b)
	fmt.Printf("%d >= %d : %t\n", a, b, a >= b)
	fmt.Printf("%d <= %d : %t\n", a, b, a <= b)
}