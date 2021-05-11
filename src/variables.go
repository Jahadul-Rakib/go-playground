package main

import (
	"fmt"
	"reflect"
)

func main() {

	var name = "Rakib"

	age := 30

	var phone string = "0180000111203"

	const COUNTRY = "Bangladesh"

	const INDIA string = "India"

	println("name: ", name, " Age: ", age, " Phone: ", phone)
	println("Name Type", reflect.TypeOf(name))
	println("Age Type: ", reflect.TypeOf(age))
	println("Phone Type: ", reflect.TypeOf(phone))

	var preDeclear string
	var preDeclear1 string

	preDeclear = "preDeclear word"
	preDeclear1 = "preDeclear1 word"

	fmt.Println(preDeclear, preDeclear1)
}
