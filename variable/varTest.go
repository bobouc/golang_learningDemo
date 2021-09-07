package main

import "fmt"

func main() {
	fmt.Println("variable test")

	var numA int
	numA = 10
	fmt.Println(numA)

	var numB int = 20
	fmt.Println(numB)

	var name = "cyber"
	fmt.Println(name)

	cyber := "good boy"
	fmt.Println(cyber)

	var m float64 //m: 0.0
	fmt.Println(m)
	var s string //""
	fmt.Println(s)
	var s1 []int
	fmt.Println(s1)
	fmt.Println(s1 == nil)

	//constant variable
	const PATH string = "http://www.google.com"
	const PI = 3.1415925535
	fmt.Println(PATH)
	fmt.Println(PI)
	// define a group of constant variable
	const c1, c2, c3 = 100, 300, "test"
	const (
		MALE   = 0
		FEMALE = 1
		UNKNOW = 3
	)
	const (
		a int = 100
		b
		c string = "const string"
		d
	)
	fmt.Printf("%T, %d\n", a, a)
	fmt.Printf("%T, %d\n", b, b)
	fmt.Printf("%T, %s\n", c, c)
	fmt.Printf("%T, %s\n", d, d)

	//enumeration 
	const(
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)
	

}
