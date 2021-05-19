package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {

	//variable declaration
	var name = "Rakib"
	var (
		AA int
		BB bool
		CC string
	)
	var (
		AAA int    = 12
		BBB bool   = true
		CCC string = "Hello"
	)
	var (
		HOME   = os.Getenv("HOME")
		USER   = os.Getenv("USER")
		GOROOT = os.Getenv("GOROOT")
	)
	fmt.Println(AA, BB, CC)
	fmt.Println(AAA, BBB, CCC)
	fmt.Println("------------", HOME, "------------")
	fmt.Println("------------", USER, "------------")
	fmt.Println("------------", GOROOT, "------------")

	//short hand declaration
	age := 30

	//variable declaration with type
	var phone string = "0180000111203"

	// declare constant variable
	const COUNTRY = "Bangladesh"
	const INDIA string = "India"
	const saturday, sunday, monday = 1, 2, 3
	const (
		male   = 0
		female = 1
	)
	const (
		jack, king, moon    = "ABC", "MOD", "CLue"
		jack1, king1, moon1 = "ABC1", "MOD1", "CLue1"
	)
	const (
		a = iota
		b
		c
	)

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
