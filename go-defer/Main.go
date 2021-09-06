package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	defer fmt.Println("middle") // exec before exit (LIFO)
	fmt.Println("end")

	fmt.Println("--------------------")
	/*
		res, err := http.Get("http://www.google.com/robots.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close() // open and close together

		robots, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", robots)
		fmt.Println("--------------------")

		http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("Hello Go!"))
		})

		err2 := http.ListenAndServe(":8081", nil)
		if err2 != nil {
			panic(err2.Error()) // I define if I panic or not
		}
	*/
	fmt.Println("--------------------")

	/*
		fmt.Println("start")
		defer func() { // anonymous function
			if err := recover(); err != nil {
				log.Println("Error", err)
			}
		}()
		panic("something bad happened")
		fmt.Println("end")
	*/
	fmt.Println("--------------------")

	fmt.Println("start")
	panicker()
	fmt.Println("end")

}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			// panic(err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking") // never executes, but end on function above does
}
