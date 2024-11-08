package main

import "fmt"

var currentYear = 2024

const CURRENT_MONTH = 11

func main() {

}

func variables() {
	// fmt.Println("Hello World")

	// var nazwaZmiennej [typ] = wartość
	// lub
	// nazwaZmiennej := wartość (tylko na poziomie funkcji)

	var firstName string = "Jan"
	var lastName = "Kowalski"
	var age int
	age = 32

	email := "jan@training.pl"

	fmt.Println(firstName + " " + lastName + " " + email)
	fmt.Println(age)

	/*
		var a, b, c int = 1, 2, 3
		var d, text = 4, "Hello"

		e, otherText := 5, "World"
	*/
}
