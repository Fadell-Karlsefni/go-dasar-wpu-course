package main

import "fmt"

func clean() {
	fmt.Println("Cleanup : Menutup Semua Resource....") // akan muncul ketika defer yang kedua sudah di eksekusi
}

func bacaConfig(namaFile string) {
	//defer akan selalu di panggil walaupun ada panic
	defer clean()

	defer func()  {
		if r := recover(); r != nil {
			fmt.Println("Terjadi error",r) // defer yang paling bawah akan di eksekusi duluan
		}
	}()
	if namaFile == "" {
		panic("Nama file configuration tidak boleh kosong")
	}

	fmt.Println("Membaca konfigurasi dari ",namaFile)
}

func main() {
	bacaConfig("")
	fmt.Println("Program tetap berjalan setelah panic") // di esekusi terakhir karena berada di bawah function bacaConfig
}