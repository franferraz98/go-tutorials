package main

import "fmt"

func main() {
	statePopulations := map[string]int{
		"California": 30000000,
		"Texas":      25000000,
	}

	if pop, ok := statePopulations["California"]; ok {
		fmt.Println(pop)
	}

	fmt.Println("--------------------------")

	number := 50
	guess := -5
	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("The number must be between 1 and 100")
	} else {
		if guess < number {
			fmt.Println("Too low")
		} else if guess > number {
			fmt.Println("Too high")
		} else {
			fmt.Println("You got it!")
		}
	}
	fmt.Println("--------------------------")

	switch i := 2 + 3; i { // implicit break
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("default")
	}

	i := 10
	switch {
	case i <= 10:
		fmt.Println("less or equal to ten")
		fallthrough // anti-break
	case i <= 20:
		fmt.Println("less or equal to twenty")
	default:
		fmt.Println("not implemented")
	}
	fmt.Println("--------------------------")

	var w interface{} = [3]int{}
	switch w.(type) { // type switch
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("Another stuff")
	}
}

func returnTrue() bool {
	fmt.Println("returning true...")
	return true
}
