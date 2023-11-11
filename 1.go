package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello World")
	//Declaration of the variable
	var number int = 78
	//Auto type detection by the compiler
	num := 56
	fmt.Println(number)
	fmt.Println(num)

	var floatNum float64 = 12345678.9
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	fmt.Println(len(myString))
	//This is used to determine the length of the special character outside the ASCII table
	fmt.Println(utf8.RuneCountInString("y"))

	//the default type of all int and float is 0
	//the default type of string is an empty string ""
	//the default type of the bool is false

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 rune
	fmt.Println(intNum3)

	var myVar string = "text"
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "const value"
	fmt.Println(myConst)

	const pi float32 = 3.1415
	fmt.Println(pi)
}
