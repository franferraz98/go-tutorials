package main

import "fmt"

func main() {
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 { // loop two variables
		fmt.Println(i, j)
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println(i)

	i = 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	i = 0 // avoid this iteration
	for i < 10 {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("--------------------")

Loop: // break outer loop
	for i = 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	fmt.Println("--------------------")

	s := []int{1, 2, 3}
	for k, v := range s { // iterate slices/arrays/strings...
		fmt.Println(k, v)
	}

	for _, v := range s { // value only
		fmt.Println(v)
	}
}
