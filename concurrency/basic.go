package main

import (
	"fmt"
	"sync"
	//"time"
)

// func main() {

// 	fmt.Println("hello i am here")

// 	go func ()  {
// 		fmt.Println("i am a goroutines")
// 	} ()

// 	time.Sleep(1 * time.Second)
// }

//safer alternative
func main() {
	var wg sync.WaitGroup // 1️⃣ Create a wait group
	wg.Add(1)             // 2️⃣ Tell it: “We're going to wait for 1 goroutine”

	go func() {
		defer wg.Done() // 4️⃣ When goroutine is done, mark it as done
		fmt.Println("i am a goroutines")
		
	}()

	fmt.Println("hello i am here")
	wg.Wait() // 3️⃣ Wait here until the goroutine signals it's done
}

/*
	❓ Why is using sync.WaitGroup safer than time.Sleep?
	🧪 Problem with time.Sleep:
		You're guessing how long the goroutine will take to finish.

		What if the goroutine takes longer than 1 second?

		Or finishes in 10 milliseconds and you’re wasting time sleeping unnecessarily?

		It doesn’t scale when you have multiple goroutines or more complex logic.

	❌ time.Sleep = unreliable, hard to manage, not scalable.

	✅ Why sync.WaitGroup is better:
		It's synchronous: It waits exactly until all goroutines finish.

		No guessing or delays — clean, predictable, and safe.

		Works for any number of goroutines.
*/