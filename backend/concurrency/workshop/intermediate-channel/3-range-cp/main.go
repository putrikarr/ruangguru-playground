package main

func squareWorker(input <-chan int, output chan<- int) {
	for i := range input {
		output <- i * i
	}
	//lakukan for range loop
	// TODO: answer here
}

func receiver(output chan<- int) {
	input := make(chan int) // mengirim 0-10 ke squareWorker
	go squareWorker(input, output)
	for i := 0; i < 10; i++ {
		input <- i
	}
	close(input)
}

//kirim nilai i ke channel
// TODO: answer here
