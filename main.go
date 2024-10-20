package main

import (
	"fmt"
	// "math/cmplx"
)

func main() {
	var (
		name string = "Robert"
		// sex bool = true

		// i8 int8 = 127
		// u8 uint8 = 255
		// i16 int16 = 32767
		// u16 uint16 = 65535
		// i32 int32 = 2147483647
		// u32 uint32 = 4294967295
		// i64 int64 = 9223372036854775807
		// u64 uint64 = 18446744073709551615

		// age int = 30 // depends on your OS

		// pi float32 = 3.14159
		// e float64 = 2.718281828459045

		// b byte = 255 // alias for uint8
		// r rune = 'a' // alias for int32

		// z1 complex64 = 1 + 2i
		// z2 complex128 = 3 + 4i
	)

	fmt.Printf("Type: %T, value: %v\n", name, name)
}
