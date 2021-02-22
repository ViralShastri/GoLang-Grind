package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("goRoutines:", runtime.NumGoroutine())

	var counter int64

	const goRoutines = 100
	var wg sync.WaitGroup
	wg.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter:\t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("goRoutines:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Count:", counter)
}
