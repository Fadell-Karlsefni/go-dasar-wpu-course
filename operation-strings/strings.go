package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Halo Dunia"

	fmt.Println("=== String Operations ===")
	fmt.Printf("Asli       : %s\n", text)

	// Ubah ke huruf kecil
	fmt.Printf("Lowercase  : %s\n", strings.ToLower(text))

	// Ubah ke huruf besar
	fmt.Printf("Uppercase  : %s\n", strings.ToUpper(text))

	// Hitung panjang string
	fmt.Printf("Length     : %d\n", len(text))

	// Ganti kata
	fmt.Printf("Replace    : %s\n", strings.ReplaceAll(text, "Halo", "Hai"))

	// Cek apakah mengandung kata tertentu
	fmt.Printf("Contains 'Halo'? %t\n", strings.Contains(text, "Halo"))
	fmt.Printf("Contains 'Hello'? %t\n", strings.Contains(text, "Hello"))

	// Split string jadi slice
	words := strings.Split(text, " ")
	fmt.Printf("Split      : %v\n", words)

	// Gabungkan slice jadi string
	joined := strings.Join(words, "-")
	fmt.Printf("Join       : %s\n", joined)
}
