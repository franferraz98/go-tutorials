package main

import "fmt"

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 5; i++ {
		fmt.Println(inc.Increment())
	}

	fmt.Println("------------------")

	var wc WriterCloser = &myWriterCloser{} // if value, only evaluates value-receiver methods. if pointer, evaluates both
	fmt.Println(wc)

}

type Writer interface { // interfaces describe behaviours, not data
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type myWriterCloser struct{}

func (mwc *myWriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

func (mwc myWriterCloser) Close() error {
	return nil
}
