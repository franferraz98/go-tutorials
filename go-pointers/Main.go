package main

import "fmt"

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, *b)
	*b = 14
	fmt.Println(a, *b)

	fmt.Println("--------------------")

	a1 := [3]int{1, 2, 3}
	b1 := &a1[0]
	c1 := &a1[1]
	fmt.Printf("%v %p %p\n", a1, b1, c1)

	fmt.Println("--------------------")

	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	ms = new(myStruct)
	(*ms).foo = 42
	fmt.Println(ms.foo) // able to access the struct params through pointer

	fmt.Println("--------------------")

	a2 := map[string]string{"foo": "bar", "baz": "buz"} // maps and slices are pointers
	b2 := a2
	fmt.Println(a2, b2)
	a2["foo"] = "qux"
	fmt.Println(a2, b2)

}

type myStruct struct {
	foo int
}
