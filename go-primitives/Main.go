package main

import "fmt"

func main() {
	grades := [...]int{97, 85, 93}
	var students [3]string
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("Students :%v\n", students)
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Number of students: %v\n", len(students))
	fmt.Println("----------------------")

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}

	fmt.Printf("Identity matrix: %v\n", identityMatrix)
	fmt.Println("----------------------")

	a := [...]int{1, 2, 3}
	b := &a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("----------------------")

	c := []int{1, 2, 3}
	fmt.Printf("Lenghth: %v\n", len(c))
	fmt.Printf("Capacity: %v\n", cap(c))
	fmt.Println("----------------------")

	x := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	y := x[:]   // all elements
	z := x[3:]  // 4th to end
	u := x[:6]  // first to 6th
	v := x[3:6] // 4th to 6th
	x[5] = 42
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(u)
	fmt.Println(v)

	y = append(y, 11) // append
	fmt.Println(y)

	y = append(y, []int{1, 2, 3}...) // append array
	fmt.Println(y)

	z = append(y[:2], y[3:]...) // remove third element *WORKS WRONG WHEN POINTING TO SAME ARRAY*
	fmt.Println(z)

	w := make([]int, 10, 100)
	fmt.Println(w)

}
