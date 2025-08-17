package main

import "fmt"

func main() {

	fmt.Println("First loop")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("second loop")
	for j := 0; j < 3; j++{
		fmt.Println(j)
	}

	fmt.Println("Third loop")
	for i := range 3 {
		fmt.Println("range", i)
	}

	fmt.Println("Fourth loop")
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("fiveth loop")
	for n := range 6 {
		if n % 2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}
