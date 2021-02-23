package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	goRotuines := 100
	var wg sync.WaitGroup
	wg.Add(goRotuines)
	for i := 0; i < goRotuines; i++ {
		go func() {
			value := counter
			runtime.Gosched()
			value++
			counter = value
			fmt.Println(counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("End Value:", counter)
}
