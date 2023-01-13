package main

import "fmt"

func if_else() {
	quiz1, quiz2 := 80, 90
	if quiz1 > quiz2 {
		fmt.Println("quiz1 is greater than quiz2")
	} else if quiz1 < quiz2 {
		fmt.Println("quiz1 is less than quiz2")
	} else {
		fmt.Println("quiz1 is equal to quiz2")
	}
}
