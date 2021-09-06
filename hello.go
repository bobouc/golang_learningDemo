package main

import "fmt"

func main(){
	var a, b *int
	numA := 10
	a = &numA
	b = &numA

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(numA)
	fmt.Println("hello")
	fmt.Println("hello")
}
