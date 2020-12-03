package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)
func main() {
	/**
	Part 1
	 */
	fmt.Println(runSlope(3,1))
	/**
	Part 2
	*/
	var slopes = [5][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	i := 1
	for _, slope := range slopes {
		i *= runSlope(slope[0],slope[1])

	}
	fmt.Println(i)

}

func runSlope(right int, down int)int {
	input,_ := ioutil.ReadFile("input.txt")
	var inputMap = strings.Split(string(input), "\n")
	var treeCount,col int


	for i :=0; i < len(inputMap)-1; i+= down {
		if []rune(inputMap[i])[col] == '#' {
				treeCount++
			}
			col += right
			col %= len(inputMap[i])
	}
	return treeCount
	}
