package main

import "fmt"

func main() {
	// signed integer
	i8 := int8(127)
	i16 := int16(32767)
	i32 := int32(2147483647)
	i64 := int64(9223372036854775807)
	i := -100

	// unsigned integer
	u8 := uint8(255)
	u16 := uint16(65535)
	u32 := uint32(4294967295)
	u64 := uint64(18446744073709551615)
	u := uint(100)

	fmt.Println("=== Signed Integer (i) ===")
	fmt.Printf("i8  : %d\n", i8)
	fmt.Printf("i16 : %d\n", i16)
	fmt.Printf("i32 : %d\n", i32)
	fmt.Printf("i64 : %d\n", i64)
	fmt.Printf("i   : %d\n", i)

	fmt.Println("\n=== Unsigned Integer (u) ===")
	fmt.Printf("u8  : %d\n", u8)
	fmt.Printf("u16 : %d\n", u16)
	fmt.Printf("u32 : %d\n", u32)
	fmt.Printf("u64 : %d\n", u64)
	fmt.Printf("u   : %d\n", u)
}
