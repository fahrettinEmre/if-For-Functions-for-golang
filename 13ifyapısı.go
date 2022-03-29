package main

import (
	"fmt"
)

func main() {
	x := 10
	if x := -5; x < 0 { // buradaki x değişkeni if yapısına ait. // 2 statement var ise ; kullanılır.
		fmt.Println(x, "negatif sayıdır.")
	} else if x%2 == 0 {
		fmt.Println(x, "çift sayıdır.")
	} else {
		fmt.Println(x, "tek sayıdır.")
	}

	fmt.Println(x)

	a := 10
	if a < 0 {
		fmt.Println(a, "negatif sayıdır.")
	} else {
		if a%2 == 0 {
			fmt.Println(a, "çift sayıdır.")
		} else {
			fmt.Println(a, "tek sayıdır.")
		}
	}
}
