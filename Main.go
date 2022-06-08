package main

import (
	"fmt"
	"strconv"
)

//at global level we can also define the constant
const (
	pie  float64 = 3.14
	User string  = "sanyam" //we can move user to the another package also
)

var globalScope = 0
var globalScope2 = 0

//if we make a global variable with pascale casing than it can be use outside the scope also
var Hello = 23

func main() {
	var i int
	i = 12
	j := "sanyam"
	var k float64 = 1.02
	s := 1.23
	var hell string = strconv.Itoa(i) // for conversion to string we have to use strconv package
	var globalScope2 = 23             // we can  reasign the value of global scope because its scope get change ow also it will work fine
	// i:=12 this give error because this cant be reasign `i`

	var x float32 = float32(i)
	//this is type casted
	fmt.Println("Hello")
	fmt.Printf("%v,%T\n", i, i)
	fmt.Printf("%v,%T\n", j, j)
	fmt.Printf("%v,%T\n", k, k)
	fmt.Printf("%v,%T\n", s, s)
	fmt.Printf("%v,%T\n", globalScope, globalScope)
	fmt.Printf("%v,%T\n", globalScope2, globalScope2)
	fmt.Printf("%v,%T\n", x, x)
	fmt.Printf("%v,%T\n", hell, hell)

	//data tyoes in go
	//bool
	var b bool = true
	fmt.Printf("%v,%T\n", b, b)
	// uint8		unsigned 8-bit integers (0 to 255)
	// uint16		unsigned 16-bit integers (0 to 65535)
	// uint32		unsigned 32-bit integers (0 to 4294967295)
	// uint64		unsigned 64-bit integers (0 to 18446744073709551615)
	// int8			signed 8-bit integers (-128 to 127)
	// int16		signed 16-bit integers (-32768 to 32767)
	// int32		signed 32-bit integers (-2147483648 to 2147483647)
	// int64		signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	//same as float
	//string

	//comlex values are also store in go
	//complex64
	//complex128
	var com complex64 = 1 + 2i
	fmt.Printf("%v,%T\n", com, com)
	com = complex(3, 4)
	fmt.Printf("%v,%T\n", real(com), real(com))
	fmt.Printf("%v,%T\n", imag(com), imag(com))
	//operators
	//arithmetic operators
	// +	Addition
	// -	Subtraction
	// *	Multiplication
	// /	Division
	// %	Modulus
	// ++	Increment
	// --	Decrement
	// **	Exponentiation
	// &	Bitwise AND
	// |	Bitwise OR
	// ^	Bitwise XOR
	// <<	Left shift
	// >>	Right shift
	// ==	Equal
	// !=	Not equal
	// >	Greater than
	// <	Less than
	// >=	Greater than or equal to
	// <=	Less than or equal to
	// &&	Logical AND
	// ||	Logical OR
	// !	Logical NOT
	// =	Assignment
	// +=	Add and assign
	// -=	Subtract and assign
	// *=	Multiply and assign
	// /=	Divide and assign
	// %=	Modulus and assign
	// &=	Bitwise AND and assign
	// |=	Bitwise OR and assign
	// ^=	Bitwise XOR and assign
	// <<=	Left shift and assign
	// >>=	Right shift and assign
	// &^=	Bitwise AND NOT and assign

	//string
	//access by str[0] but we cant change value if we want to change value we need to convert in array and than change the value
	var str = "hello world"
	fmt.Printf("%v,%T\n", str[0], str[0])
	fmt.Printf("%v,%T\n", string(str[0]), str[0])
	//change the vale
	str1 := []byte(str)
	fmt.Printf("%v,%T\n", str1, str1)
	str1[0] = 'S'
	fmt.Printf("%v\n", string(str1))

	//const is the keyword for constant
	const y int = 12
	fmt.Println(y)
	//we can also define the constant at global scope
	fmt.Println(pie + k)
	fmt.Println(User)
	const (
		one   = iota
		two   = iota
		three = iota
	)
	fmt.Println(one, two, three)
	//because in this scope the value is increasing
	//we can also write in this way
	const (
		one1 = iota
		two2
		three3
		three4
	)
	fmt.Println(one1, two2, three3, three4)
	//if we want to leave any of the incremenet we will use underscore _
	// example
	const (
		one2 = iota
		_
		_
		two1
		three1
		_ //in this way we can leave one increment
		three2
	)
	fmt.Println(one2, two1, three1, three2)

	//array
	var arr [8]int = [8]int{1, 2,4,5,6,7,8,9};
	fmt.Println(arr);
	fmt.Println(len(arr));//length of array
	fmt.Println(cap(arr));//capacity of arary;
	var arr1[3]int = [3]int{};
	arr1[0] = 3;
	amount:=arr1// this line copy the array to another array and not just point to that 
	amount[0] = 1
	fmt.Println(arr1);
	fmt.Println(amount);

	//if we dont want to copy we just want to point
	arr2 := &amount
	arr2[0] = 2
	fmt.Println(arr2);
	fmt.Println(amount);
	arr3 := arr[:]
	b1 := arr[:len(arr)-1]
	c1 := arr[2:]
	d1 := arr[:5]
	e1 := arr[2:7]
	fmt.Println(arr3);
	fmt.Println(b1);
	fmt.Println(c1);
	fmt.Println(d1);
	fmt.Println(e1);


	

}
