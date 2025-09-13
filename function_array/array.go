package main

import "fmt"

func main() {
	angka := [5]int{1, 2, 3, 4, 5}

	// mengetahui nilai index array
	fmt.Println("<=== CEK SALAH SATU NILAI INDEX ===>")
	fmt.Println("Jumlah element : ", len(angka))
	fmt.Println("Index ke 1 : ", angka[1])

	// mengubah nilai index yang ada di array
	fmt.Println("<=== MENGUBAH NILAI ARRAY ===>")
	angka[4] = 6
	fmt.Println("Index ke 4 : ",angka[4])
	fmt.Println(angka)

	// perulangan atau looping
	fmt.Println("<=== INI ADALAH PERULANGAN ===>")
	for index, value := range angka {
		fmt.Println("Isi Index ke : ",index," = ",value)
	}

	// perbandingan dalam array
	fmt.Println("<=== INI ADALAH PERBANDINGAN ===>")
	arr1 := [3]int {1,2,3}
	arr2 := [3]int {2,4,6}

	fmt.Println("Apakah array nya sama : ",arr1==arr2)
	fmt.Println("Apakah array1 Tidak sama dengan Array2 : ",arr1!=arr2)

}