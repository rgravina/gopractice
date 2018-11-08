package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		printIteration(i)
	}

	i := 0
	for {
		if i >= 10 {
			break
		}
		printIteration(i)
		i++
	}
}

func printIteration(i int) (int, error) {
	return fmt.Println("Loop iteration: ", i+1)
}
