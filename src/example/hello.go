package main

import (
	"errors"
	"fmt"
	"math"
)

//struct is a collection of fields group thing together to create more
//logical type name
type person struct {
	//each field name and type
	name string
	age  int
}

func main() {
	//abstraction slices
	fmt.Println("Hello, World!")
	a := []int{5, 4, 3, 2, 1}
	//return an new slice
	a = append(a, 13)
	fmt.Println(a)
	//maps like dict in python or associate array in php
	//[]- types of the keys, than types of the values in the map
	// create a map
	vertices := make(map[string]int)
	//indexing arrays

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12
	//get the value for the particular key
	fmt.Println(vertices["triangle"])
	//delete function to remove something from the map
	delete(vertices, "squares")
	//the only type of loop is for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//while  loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
	//iterate over each element or a slice by using range
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}
	//same thing for the map
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key:", key, "value", value)
	}
	//sum call
	result := sum(2, 3)
	fmt.Println(result)
	//sqrt  call
	resultSqrt, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultSqrt)
	}
	// to create struct of the type
	//set the fields as initialize array
	p := person{name: "Vadim", age: 23}
	fmt.Println(p.age)

	//pointers in go
	k := 7
	//k is copied by value, so after inc func called result descarded
	inc(k)
	//if pass the poiter to k the funcition would be able to modify the
	//original version
	inc1(&k)
	fmt.Println(k)
	//gives pointer to k
	fmt.Println(&k)
}

// create new function (name type)return value
func sum(x int, y int) int {
	return x + y
}

// multiple arguments, error built in type, doesn't have exceptions
func sqrt(x float64) (float64, error) {
	//return error if negative
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

//function that increments the variable
func inc(x int) {
	x++
}

// accept the pointer with x *int
// also check if go accepts polymorhism, only by using interfaces
func inc1(x *int) {
	//dereference the pointer
	*x++
}
