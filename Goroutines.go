package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
//when we use go we use wait group that it will wait for the counter to go to 0
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		//go keyword will make it concurrent that means it will not call it one after completion of other
		go dbCall(i)
	}
	fmt.Printf("\nTotal execution tim : %v", time.Since((t0)))
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is :*", dbData[i])
	wg.Done()
}
