package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Print("Size of spiral: ")
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		fmt.Println(err)
	}

	if size%2 == 0 {
		size++
	}

	for _, row := range GenerateUlamSpiral(size) {
		for _, v := range row {
			fmt.Printf("%d\t", v)
		}
		fmt.Print("\n")
	}
}

func GenerateUlamSpiral(n int) [][]int {
	spiral := make([][]int, n)

	for k, _ := range spiral {
		spiral[k] = make([]int, n)
	}
	dir := 3 // directions are RIGHT = 3, UP = 12, DOWN = 6, LEFT = 9
	mid := n / 2

	total := n * n
	x, y := mid, mid

	for i := 1; i <= total; i++ {
		spiral[x][y] = i

		/* Looking for solution
		if i == 361527 {
			fmt.Printf("Solution: %d\n\n", y-x)
			return spiral
		}
		*/

		switch dir {
		case 3:
			if (y+1 < n) && (spiral[x][y+1] == 0) {
				dir = 12
				y++
			} else {
				x++
			}
		case 12:
			if (x-1 >= 0) && (spiral[x-1][y] == 0) {
				dir = 9
				x--
			} else {
				y++
			}
		case 9:
			if (y-1 >= 0) && (spiral[x][y-1] == 0) {
				dir = 6
				y--
			} else {
				x--
			}
		case 6:
			if (x+1 < n) && (spiral[x+1][y] == 0) {
				dir = 3
				x++
			} else {
				y--
			}
		}
	}
	return spiral
}
