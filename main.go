package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	// uint8  : 0 to 255
	// uint16 : 0 to 65535
	// uint32 : 0 to 4294967295
	// uint64 : 0 to 18446744073709551615
	// int8   : -128 to 127
	// int16  : -32768 to 32767
	// int32  : -2147483648 to 2147483647
	// int64  : -9223372036854775808 to 9223372036854775807
	// int    : -9223372036854775808 to 9223372036854775807

	var i32 int32
	var i64 int64
	var i int

	// unsafe size
	fmt.Println("Unsafe Size int32")
	fmt.Println(unsafe.Sizeof(i32))
	fmt.Println("Unsafe Size int64")
	fmt.Println(unsafe.Sizeof(i64))
	fmt.Println("Unsafe Size int")
	fmt.Println(unsafe.Sizeof(i))
	
	const MaxUint = ^uint(0)
	const MinUint = 0
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1
	
	const MaxInt32 = math.MaxInt32
	const MinInt32 = math.MinInt32
	const MaxInt64 = math.MaxInt64
	const MinInt64 = math.MinInt64

	// Max Min Digit	
	fmt.Println(" \n ")
	fmt.Println("int32")
	fmt.Println(MaxInt32)
	fmt.Println(MinInt32)
	
	fmt.Println(" \n ")
	fmt.Println("int64")
	fmt.Println(MaxInt64)
	fmt.Println(MinInt64)

	fmt.Println(" \n ")
	fmt.Println("int")
	fmt.Println(MaxInt)
	fmt.Println(MinInt)
	
	// Compare digit
	fmt.Println(" \n ")
	fmt.Println("MaxInt == MaxInt64")
	fmt.Println(MaxInt == MaxInt64)
}
