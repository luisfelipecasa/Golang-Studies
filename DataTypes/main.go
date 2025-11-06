package main

import "fmt"

func main() {

	// Signed Integers
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	i = -128
	i8 = 127
	i16 = -32224
	i32 = 298374
	i64 = -38247982

	// Unsigned Integers
	var u uint
	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64

	u = 128
	u8 = 127
	u16 = 32224
	u32 = 298374
	u64 = 38247982

	fmt.Println("Signed integers:", i, i8, i16, i32, i64)
	fmt.Println("Unsigned integers:", u, u8, u16, u32, u64)

	// Floaring point
	var f32 float32 = 10.6
	var f64 float64 = 10.6

	fmt.Println("Float32:", f32)
	fmt.Println("Float64:", f64)

	//diferença de precisão entre os floats
	var HP float64 = 10123456789012345
	var LP float32 = 10123456789012345

	fmt.Println("menos precisão Float32:", LP)
	fmt.Println("maior precisão Float64:", HP)

	//Boolean
	var isActive bool = true
	var isOn bool = false

	fmt.Println("esta ativo:", isActive)
	fmt.Println("esta ligado:", isOn)

	//Complex Data
	var CN1 complex128 = complex(5, 10)
	var CN2 complex64 = complex(2, 7)

	fmt.Println("CN1:", CN1)
	fmt.Println("CN2:", CN2)

	// String
	var name string = "Rafael Jaques"

	fmt.Println("Meu professor é:", name)

}
