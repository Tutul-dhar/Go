package main

import "fmt" 

// func main() {
// 	ch := make(chan string) // 1️⃣ create a channel

// 	go func() {
// 		ch <- "Hello from goroutine!" // 2️⃣ send to channel
// 	}()

// 	msg := <-ch // 3️⃣ receive from channel
// 	fmt.Println(msg)
// }

// buffered channels
// func main() {
// 	ch := make(chan int, 2) // channel with capacity 2

// 	ch <- 1
// 	ch <- 2
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

func digit(number int, digits chan int) {
	for number != 0 {
		digits <- number%10
		number /= 10
	}

	close(digits)
}

func calsqr(number int, sqrch chan int) {
	digits := make(chan int)
	go digit(number,digits)
	sum := 0
	for digit := range digits {
		fmt.Println(digit)
		sum += digit*digit
	}

	sqrch <- sum
}

func main() {
	number := 456

	sqrch := make(chan int)

	go calsqr(number, sqrch)

	fmt.Println(<-sqrch)
}