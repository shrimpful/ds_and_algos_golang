package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var threedArray [2][2][2]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				threedArray[i][j][k] = rand.Intn(3)
			}
		}
	}
	fmt.Println(threedArray)
}
