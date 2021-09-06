package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	name   string `required max:"100"` // prequisite called tag
	origin string
}

type Bird struct {
	Animal   // inheritance equivalent, it's called composition
	speedKPH float32
	canFly   bool
}

type Doctor struct {
	number     int
	actorName  string
	companions []string
	episodes   []string
}

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 39000000,
		"Texas":      27000000,
		"Ohio":       11000000,
	}

	statePopulations["Georgia"] = 10000000 // not ordered
	delete(statePopulations, "California")

	pop, ok := statePopulations["Oho"] // check existance

	fmt.Println(statePopulations, pop, ok, len(statePopulations))
	fmt.Println("-------------------")
	sp := statePopulations
	delete(sp, "Ohio") // deletes from both maps
	fmt.Println(statePopulations)
	fmt.Println("-------------------")

	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}
	fmt.Println(aDoctor.actorName, aDoctor.number, aDoctor.companions[1])
	fmt.Println("-------------------")

	aDoctor2 := struct{ name string }{name: "John Pertwee"} // anonymous struct, not recommended
	anotherDoctor := &aDoctor2
	anotherDoctor.name = "Tom Baker"
	fmt.Println(aDoctor2.name)
	fmt.Println("-------------------")

	b := Bird{}
	b.name = "Emu"
	b.origin = "Australia"
	b.speedKPH = 48.0
	b.canFly = false
	fmt.Println(b)

	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)

}
