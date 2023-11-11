package main

import "fmt"

func main() {
	//Different way of initialising the array in Golang

	//  var intArr [3]int32
	// var intArr [3]int32 = [3]int32{1, 2, 3}
	// intArr := [3]int32{1, 2, 3}
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)
	intArr[1] = 123
	fmt.Println(intArr[0])
	//It is slicing starting from 1 and going to 3 but not including it
	fmt.Println(intArr[1:3])

	fmt.Println(intArr[0])
	fmt.Println(intArr[1])
	fmt.Println(intArr[2])

	//They are stored in continuous location after every 4 bits
	//& is used to access the address or the location
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

}
