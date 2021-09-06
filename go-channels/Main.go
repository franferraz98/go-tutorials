package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var doneCh = make(chan struct{})
var logCh = make(chan int)

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) { // Read only
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) { // Write only
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()

	fmt.Println("-----------------")

	wg.Add(2)
	ch = make(chan int, 50)
	go func(ch <-chan int) { // Continuous listen
		for i := range ch {
			fmt.Println(i)
		}

		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch) // Close channel (mandatory)
		wg.Done()
	}(ch)
	wg.Wait()

	fmt.Println("-----------------")

	ch = make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) { // Continuous listen
		for {
			if i, ok := <-ch; ok { // able to check ok
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch) // Close channel (mandatory)
		wg.Done()
	}(ch)
	wg.Wait()

	go logger()
	logCh <- 55
	doneCh <- struct{}{} // Signal only
}

func logger() {
Loop:
	for {
		select {
		case entry := <-logCh:
			fmt.Println("Entry:", entry)
		case <-doneCh: // End case
			break Loop
		}
	}
}
