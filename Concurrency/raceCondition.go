package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("goRoutines:", runtime.NumGoroutine())

	counter := 0

	const goRoutines = 100
	var wg sync.WaitGroup
	wg.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func() { v := counter; runtime.Gosched(); v++; counter = v; wg.Done() }()
		fmt.Println("goRoutines:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Count:", counter)
}
