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

	//Slicing in Array

	var IntSlice []int32 = []int32{4, 5, 6}
	fmt.Println(IntSlice)
	fmt.Printf("The length is %v with the capacity %v", len(IntSlice), cap(IntSlice))
	fmt.Println()
	IntSlice = append(IntSlice, 7)
	fmt.Println(IntSlice)
	fmt.Printf("The length is %v and the capacity is %v \n", len(IntSlice), cap(IntSlice))

	var intSlice2 []int32 = []int32{8, 9}
	IntSlice = append(IntSlice, intSlice2...)
	fmt.Println(IntSlice)

	//Maps it is a key value pairs

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["Jason"]
	delete(myMap2,"Adam")
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

}
