package main

import "fmt"

func main() {
	num := 10
	if num > 10 {
		fmt.Println("bigger than 10")
	} else if num == 10 {
		fmt.Println("value equal to 10")
	} else {
		fmt.Println("smaller than 10")
	}

	/*
		using fallthrough to execute the following cases.
		but it can be stopped by break.
	*/
	switch x := 3; x {
	case 1:
		fmt.Println("x == 1")
	case 2:
		fmt.Println("x == 2")
	case 3:
		fmt.Println("x == 3")
		if x > 2 {
			break
		}
		fallthrough
	case 4:
		fmt.Println("x == 4")
	}
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	/*
		equal to while in c/cpp
	*/
	// index := 1
	// for index == 1 {
	// 	fmt.Println("looping")
	// }

	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [4]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

}
