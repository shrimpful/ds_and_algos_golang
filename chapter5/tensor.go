package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var array [3][3][3]int
	var i int
	var j int
	var k int
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			for k = 0; k < 3; k++ {

				array[i][j][k] = rand.Intn(3)
			}
		}
	}
	fmt.Println(array)

	fmt.Println("zero mode unfold")
	for j = 0; j < 3; j++ {
		for k = 0; k < 3; k++ {
			fmt.Printf("%d ", array[0][j][k])
		}
		fmt.Printf("\n")
	}
	fmt.Println("1-mode unfold")
	for j = 0; j < 3; j++ {
		for k = 0; k < 3; k++ {
			fmt.Printf("%d ", array[1][j][k])
		}
		fmt.Printf("\n")
	}
	fmt.Println("2-mode unfold")
	for j = 0; j < 3; j++ {
		for k = 0; k < 3; k++ {
			fmt.Printf("%d ", array[2][j][k])
		}
		fmt.Printf("\n")
	}

}
