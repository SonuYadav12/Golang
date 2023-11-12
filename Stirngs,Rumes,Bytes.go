package main

import (
	"fmt"
	"strings"
)

func main() {
	//STIRNG ARE IMMUTABLE IN GO ,ONCE CREATED CANNOT BE MODIFIED
	var myString = []rune("resume")
	var indexed = myString[0]
	fmt.Printf("%v,%T", indexed, indexed)
	fmt.Println()
	fmt.Println(myString)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	//length is not the number of digit or character it is about the number of bytes
	fmt.Printf("The length of the string is %v", len(myString))

	//The rune type in Go is an alias for int32 . Given this underlying int32 type, the rune type holds a signed 32-bit integer value. However, unlike an int32 type, the integer value stored in a rune type represents a single Unicode character.

	var myRune = "a"
	fmt.Printf("\nmyRune =%v", myRune)
	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)

	//or

	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}
