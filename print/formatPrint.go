package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Println("hello world")

	a := 1;
	b := 3.1415
	c := true
	d := "hello"
	e := 'e'
	f := "F"
	fmt.Printf("%T, %d\n", a, a)
	fmt.Printf("%T, %f\n", b, b)
	fmt.Printf("%T, %t\n", c , c)
	fmt.Printf("%T, %s\n", d, d)
	fmt.Printf("%T, %c\n", e, e)
	fmt.Printf("%T, %p\n", f, &f)

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)

	/* -------------------- */
	var numA int
	var numB float32
	fmt.Scanln(&numA, &numB)
	fmt.Printf("numA: %d,numB: %f\n", numA, numB)
	fmt.Scanf("%d,%f", &numA, &numB)
	fmt.Printf("numA: %d, numB: %f\n", numA, numB )


	/* -------------------- */
	fmt.Println("Please input:")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	fmt.Println("Value of input: "+s1)

}