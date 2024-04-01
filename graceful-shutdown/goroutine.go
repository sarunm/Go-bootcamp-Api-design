package main

import (
	"fmt"
	"time"
)

func slow(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s, ":", i)

	}
}

func main() {

	done := make(chan bool)

	go func() {
		slow("hello")
		done <- true
	}()

	//go func() {
	//    slow("world")
	//    done <- true
	//}

	<-done

	//time.Sleep(10 * time.Second)
	fmt.Println("done")
}
