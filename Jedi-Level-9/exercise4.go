package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	goRotuines := 100
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(goRotuines)
	for i := 0; i < goRotuines; i++ {
		go func() {
			mu.Lock()
			counter++
			fmt.Println(counter)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("End Value:", counter)
}
