package main

import (
	"fmt"
	"time"
)

func convertToZone(t time.Time, zone string) time.Time {
	loc,err := time.LoadLocation(zone)
	if err != nil {
		fmt.Println("Error memuat lokasi", err)
		return t
	}
	return t.In(loc)
}

func main() {

	// Ambil waktu saat ini dalam utc
	nowUTC := time.Now().UTC()
	fmt.Println("Waktu UTC : ", nowUTC)

	// konversi ke zona waktu lokal
	makassarTime := convertToZone(nowUTC,"Asia/Makassar")
	fmt.Println("Waktu Banjarmasin : ",makassarTime)

	// konversi menggunakan wakttu lain misal nya menggunakan DST(misal nya NewYork)
	NewYork := convertToZone(nowUTC,"America/New_York")
	fmt.Println("Waktu NewYork : ", NewYork)

	// menyimpan waktu (simulai ke dalam database)
	fmt.Println("Di simpan ke database (UTC) : ",nowUTC)
}