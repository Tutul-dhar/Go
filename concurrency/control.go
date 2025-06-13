package main

import (
	"fmt"
	"sync"
)

func solve() {
	fmt.Println("just call once")
}

func main() {

	var wg sync.WaitGroup
	var once sync.Once
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int)  {
			defer wg.Done()
			once.Do(solve)
		}(i)
	}

	wg.Wait()
}