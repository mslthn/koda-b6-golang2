package main

import "fmt"

func main() {
	slice()

	scores := []int{50, 75, 66, 20, 32, 90}
	newNumber := 88
	index := 3
	scores = append(scores[:index], append([]int{newNumber}, scores[index:]...)...)
	
	for i := range len(scores) {
		fmt.Println(scores[i])
	}
}
