package main

import "fmt"

func main() {
	x := 10

	x = 5
	fmt.Println("Mengubah Value x : ",x)

	x += 2 // fungsi += adalah
	fmt.Println("Ini Menambahkan Value x : ",x) // menambahkan value ke variable x
	
	x -= 2 
	fmt.Println("Ini Mengurangkan value x nya : ",x)

	x *= 2
	fmt.Println("Ini Mengalikan value x : ", x)

	x /= 2
	fmt.Println("Ini Membagikan value x : ", x)

	x %= 2
	fmt.Println("ini Modulus value x : ", x)
}