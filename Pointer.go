package main

import "fmt"

func main(){
	var p *int32=new(int32)
	var i int32
	fmt.Printf("The value of p points to is : %v",*p)
	fmt.Printf("\nThe value of i is: %v",i)
}
