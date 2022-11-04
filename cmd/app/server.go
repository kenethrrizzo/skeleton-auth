package main

import "log"

func main() {
	x := 2

	for i := 0; i < 10; i++ {
		if x%2 == 0 {
			log.Println("Hello, World!")
		} else {
			log.Println("OMG! You are gay")
		}
		x++
	}

}
