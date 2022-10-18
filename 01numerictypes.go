package main

import "fmt"

/*
Signed Integers
int8 (8-bit signed integer whose range is -128 to 127)
int16 (16-bit signed integer whose range is -32768 to 32767)
int32 (32-bit signed integer whose range is -2147483648 to 2147483647)
int64 (64-bit signed integer whose range is -9223372036854775808 to 9223372036854775807)

Unsigned Integers
uint8 (8-bit unsigned integer whose range is 0 to 255 )
uint16 (16-bit unsigned integer whose range is 0 to 65535 )
uint32 (32-bit unsigned integer whose range is 0 to 4294967295 )
uint64 (64-bit unsigned integer whose range is 0 to 18446744073709551615 )
*/

func main() {
	/*
		var overflow uint8
		overflow = 267	// uint8 range is 0-255
	*/

	// Do type casting
	// Attempting to convert to a type that has
	// a lower range than your current range will cause
	// data loss or truncation to occur
	var x int32
	var y uint32
	var z uint8
	x = 26700
	y = uint32(x) // data preserved because number is within range
	z = uint8(x)  // data loss due to out of range conversion
	fmt.Println("x:", x, "y:", y, "z:", z)
}
