package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)

	var msg = "message"
	wg.Add(1)
	go func(msg string) { // decouple variable from goroutine
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Goodbye" // go run -race to check if there is a race condition
	wg.Wait()

	fmt.Println("-----------------")

	runtime.GOMAXPROCS(4) // number of OS threads

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello2()
		m.Lock()
		go increment()
	}
	wg.Wait()

}

func sayHello() {
	fmt.Println("Hello Go!")
}

func sayHello2() {
	fmt.Printf("Hello %v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
