package main

import "fmt"

func main() {
	intslice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, dt := range intslice {
		if dt%2 == 0 {
			fmt.Println(fmt.Sprint(dt) + " is even")
		} else {
			fmt.Println(fmt.Sprint(dt) + " is odd")
		}

	}

}
