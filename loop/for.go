package main

import "fmt"

func main() {
	// 1. For biasa (sudah kamu buat)
	fmt.Println("=== For loop biasa ===")
	for i := 1; i <= 5; i++ {
		fmt.Println("angka ke", i)
	}

	// 2. While style (hanya kondisi)
	fmt.Println("\n=== While style: angka genap sampai 10 ===")
	j := 2
	for j <= 10 {
		fmt.Println(j)
		j += 2
	}

	// 3. Infinite loop dengan break
	fmt.Println("\n=== Infinite loop dengan break: countdown dari 5 ===")
	k := 5
	for {
		fmt.Println(k)
		k--
		if k == 0 {
			break
		}
	}

	// 4. Range loop: iterasi elemen slice
	fmt.Println("\n=== Range loop: iterasi slice ===")
	names := []string{"Andi", "Budi", "Cici"}
	for idx, name := range names {
		fmt.Printf("Index %d: %s\n", idx, name)
	}
}
