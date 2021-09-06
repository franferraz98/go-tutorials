package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		sayMessage("Hello Go!", i)
	}

	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name)
	fmt.Println(name)

	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is", s)

	d, err := divide(5.0, 2.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	func(d float64) {
		fmt.Println("Hello Go!", d)
	}(d)

	f := func() { // define function as variable
		fmt.Println("Hello 2!")
	}

	f()

	g := greeter{
		greeting: "Hello",
		name:     "GO",
	}

	g.greet()

	fmt.Println("Name: ", g.name)

}

func sayMessage(msg string, idx int) {
	fmt.Println(msg, idx)
}

func sayGreeting(greeting, name *string) { // same type for multiple variables
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

func sum(values ...int) int { // slice of any size
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func sum2(values ...int) (result int) { // explicit return value
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() { // type method for greeter
	fmt.Println(g.greeting, g.name)
	g.name = "Tom"
}
