package main

import (
	"fmt"
	"reflect"
)

func main() {

	//variable declaration
	var name = "Rakib"

	//short hand declaration
	age := 30

	//variable declaration with type
	var phone string = "0180000111203"

	// declare constant variable
	const COUNTRY = "Bangladesh"
	const INDIA string = "India"

	println("name: ", name, " Age: ", age, " Phone: ", phone)
	println("Name Type", reflect.TypeOf(name))
	println("Age Type: ", reflect.TypeOf(age))
	println("Phone Type: ", reflect.TypeOf(phone))

	//pre define variables where must define type
	var preDeclare string
	var preDeclare1 string

	preDeclare = "preDeclare word"
	preDeclare1 = "preDeclare1 word"

	fmt.Println(preDeclare, preDeclare1)
}
