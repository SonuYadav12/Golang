package main

import "fmt"

func main() {
	var c = make(chan int)

	// Use a goroutine for sending
	// go func() {
	// 	c <- 1
	// }()

	// go process(c)
	// fmt.Println(<-c)

	// // Use the main goroutine for receiving
	// var i = <-c
	// fmt.Println(i)

	go loop(c)
	for i:=range c{
		fmt.Println(i)
	}
}

// func process(c chan int) {
// 	c <- 123
// }
func loop(c chan int){
	defer close(c)
	for i:=0;i<5;i++{
		c <- i
	}
	fmt.Println("Exiting Process")
}
