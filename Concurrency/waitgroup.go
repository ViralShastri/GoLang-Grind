package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("Operating System\t", runtime.GOOS)
	fmt.Println("Architecture\t\t", runtime.GOARCH)
	fmt.Println("Numbers of CPUs\t\t", runtime.NumCPU())
	fmt.Println("Numbers of GOROUTIINEs\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()
	fmt.Println("Numbers of GOROUTIINEs\t", runtime.NumGoroutine())
	wg.Wait()

}

func foo() {
	for i := 1; i <= 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 1; i <= 10; i++ {
		fmt.Println("bar:", i)
	}
}
