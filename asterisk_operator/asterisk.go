package main

import "fmt"

type Addres struct {
	City, Province, Country string
}

func main() {
	addres1 := Addres{"Subang", "Jawa barat", "Indonesia"}
	addres2 := &addres1 // pointer

	addres2.City = "Bandung"
	fmt.Println(addres1) // Ikut berubah
	fmt.Println(addres2) // data yang berubah, maka data awal berubah juga

	*addres2 = Addres{"Banjar","kalsel","Indonesia"}
	fmt.Println(addres1)
	fmt.Println(addres2)
}