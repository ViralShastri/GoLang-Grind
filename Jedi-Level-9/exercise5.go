package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	goRotuines := 100
	var wg sync.WaitGroup
	wg.Add(goRotuines)
	for i := 0; i < goRotuines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("End Value:", counter)
}
