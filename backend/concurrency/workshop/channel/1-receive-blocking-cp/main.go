package main

import "fmt"

func receiveBlock(output chan int) {
	c := make(chan int)
	result := 0
	go func() {
		fmt.Println("send to channel c")
		c <- 1
		// fmt.Println("send to channel output")
		// output <- 1
		//kirim 1 ke channel c
		// TODO: answer here
	}()

	//result menerima data dari channel c
	// TODO: answer here
	result = <-c
	// fmt.Println("result:", result)
	output <- result
	fmt.Println(c) //agar variabel c digunakan
}
