package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go func() { fmt.Println("goRoutine 1"); wg.Done() }()
	go func() { fmt.Println("goRoutine 2"); wg.Done() }()

	wg.Wait()

}
