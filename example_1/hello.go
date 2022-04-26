package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

func main() {

	// VARIABLES

	var x int = 5
	y := 7
	var sum int = x + y

	fmt.Println(sum)

	// IF CONDITIONALS

	if x > 6 {
		fmt.Println("x is more than 6")
	} else if x < 2 {
		fmt.Println("x is less than 2")
	} else {
		fmt.Println("x is between 2 and 6")
	}

	// ARRAY

	// fixed size arrays
	var a [5]int
	b := [5]int{1, 2, 3, 4, 5}

	// dynamic size arrays with append
	c := []int{}
	c = append(c, 2)

	fmt.Println(a)
	fmt.Println(b)

	// MAPS (DICTS)

	numeros := make(map[string]int)
	numeros["uno"] = 1
	numeros["dos"] = 2

	fmt.Println(numeros)
	delete(numeros, "dos")
	fmt.Println(numeros)

	// LOOPS

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	j := 0

	for j < 5 {
		fmt.Println(j)
		j++
	}

	// with ranges
	letras := make(map[string]int)
	letras["a"] = 0
	letras["b"] = 1
	letras["c"] = 2

	for key, value := range letras {
		fmt.Println("Key: ", key, "Value: ", value)
	}

	// STRUCTS

	p := person{name: "Hola", age: 1}
	fmt.Println(p)
	fmt.Println(p.name)

	// POINTERS

	// the & symbol is a pointer, without it the increment function
	// wont work as it will create a copy of i, so we pass the pointer of memory

	i := 7
	increment(&i)
	fmt.Println(i)
}

// FUNCTIONS

func suma(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative num")
	}

	return math.Sqrt(x), nil
}

func printMessage(msg string) {
	fmt.Println(msg)
}

// the *int means we accept a pointer as an argument
// the *variable means we are modifying the origin variable by the pointer
func increment(x *int) {
	*x++
}
