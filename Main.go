package main

import (
	"fmt"
)

func main() {
	Num := 0
	EnteredNum := 0
	Diff := 0
	for {

		fmt.Println("The current number is: ", Num)

		fmt.Println("Please enter desired number")

		fmt.Scanln(&EnteredNum)

		Diff = EnteredNum - Num

		Num = Diff + Num

	}
}
