package main

import "fmt"

func main() {
	angka := 3

	switch angka {
	case 1:
		fmt.Println("angka nya adalah satu")
	case 2:
		fmt.Println("angka nya adalah dua")
	case 3:
		fmt.Println("angka nya adalah tiga")
	default :
		fmt.Println("angka tidak ada")
	}
}