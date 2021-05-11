package main

import "reflect"

func main() {

	var name = "rakib"
	age := 30
	var phone string = "rakib"

	println("name: ", name, " Age: ", age, " Phone: ", phone)
	println("Name Type", reflect.TypeOf(name))
	println("Age Type: ", reflect.TypeOf(age))
	println("Phone Type: ", reflect.TypeOf(phone))

}
