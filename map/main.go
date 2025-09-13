package main

import "fmt"

func main() {

	// membuat map baru dengan nama mahasiswa dan untuk key nya string dan retrun value nya string
	mahasiswa := map[string]string{
		"Nama":  "Fadillah",
		"kelas": "XII 2",
	}

	fmt.Println("Nama : ", mahasiswa["Nama"]) // menampilkan value key nya dengan memangil key itu sendiri
	fmt.Println("Kelas : ", mahasiswa["kelas"])

	mahasiswa ["Jurusan"] = "Rekayasa Perangkat Lunak" // menambahkan key baru ke map
	fmt.Println("Jurusan : ", mahasiswa["Jurusan"])
}