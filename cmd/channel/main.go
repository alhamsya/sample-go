package main

import "fmt"

func main() {
	number := make(chan int)

	go func(n int) {
		number <- n
	}(9)

	result := <-number
	fmt.Println(result)
}
