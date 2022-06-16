package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz() {
	const ENDPOINT int = 15

	for i := 1; i <= ENDPOINT; i++ {
		if i%15 == 0 {
			fmt.Print(strconv.Itoa(i), " Fizzbuzz\n")
		} else if i%5 == 0 {
			fmt.Print(strconv.Itoa(i), " Buzz\n")
		} else if i%3 == 0 {
			fmt.Print(strconv.Itoa(i), " Fizz\n")
		} else {
			fmt.Println(i)
		}
	}
}
