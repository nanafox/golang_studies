package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/* a function to greet the user */
func say_hello(name string) {
	fmt.Println("Hello, " + name + "!")
}

/* a simple addition function, works for both ints and floats */
func add(x float64, y float64) float64 {
	return x + y
}

func main() {
	/* introductions */
	say_hello("theLazyProgrammer")

	var sum float64 = add(5, 10.45)
	fmt.Println(sum)

	/* Declaring and initializing arrays */
	var number[4]int
	number[0] = 56
	fmt.Println(number[0])

	/* checking the type of data */
	fmt.Println("\nChecking the data types")
	fmt.Println(reflect.TypeOf(sum))
	fmt.Println(reflect.TypeOf(say_hello))
	fmt.Println(reflect.TypeOf(number))

	/* type conversion */
	fmt.Println("\nType Conversion")
	var is_new = true
	fmt.Println(reflect.TypeOf(is_new))
	fmt.Println("is_new(bool):", is_new)
	var str string = strconv.FormatBool(is_new)
	fmt.Println("str(string): " + str)
	fmt.Println(reflect.TypeOf(str))

	/* convert a number to string and back */
	num, err := strconv.Atoi("76")
	/* check for errors before printing value */
	if (err == nil) {
		fmt.Println(num, reflect.TypeOf(num))	
	}

	/* convert the number to string */
	str_num := strconv.Itoa(num)
	fmt.Println(str_num, reflect.TypeOf(str_num))
}
