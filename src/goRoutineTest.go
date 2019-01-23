package main

import (
	"fmt"
	"sync"
)

func myFunc(wg *sync.WaitGroup){
	fmt.Println("Hello World")
	wg.Done()
}
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go myFunc(&wg)
	}

	wg.Wait()
	fmt.Println("Finished")
}
