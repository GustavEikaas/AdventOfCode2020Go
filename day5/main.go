package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func main(){
	input,_ := ioutil.ReadFile("input.txt")
	var lines = strings.Split(string(input), "\n")
	passport := make([]int, len(lines))
	for i:=0; i<len(lines); i++{
		var rowLowRange,row,col,colLowRange float64
		rowHighRange,colHighRange := 127.0, 7.0
		line := lines[i]

		for j:=0; j<len(line); j++ {
			switch []rune(line)[j] {
			case 'F':
				if rowHighRange - rowLowRange == float64(1) {
					row = rowLowRange
				}else{
					rowHighRange -= math.Round((rowHighRange - rowLowRange) / 2)
				}
				break

			case 'B':
				if rowHighRange - rowLowRange == float64(1) {
					row = rowHighRange
				}else{
					rowLowRange += math.Round((rowHighRange - rowLowRange) / 2)
				}
				break

			case 'R':
				if colHighRange - colLowRange == 1 {
					col = colHighRange
				}else{
					colLowRange += math.Round((colHighRange - colLowRange) / 2)
				}
				break

			case 'L':
				if colHighRange - colLowRange == 1 {
					col = colLowRange
				}else{
					colHighRange -= math.Round((colHighRange - colLowRange) / 2)
				}
				break
			}
		}
		seatID := int(row * 8 + col)
		passport[i] = seatID
	}

	var maxSeatID int
	for i:=0; i<len(passport); i++{
		if passport[i] > maxSeatID{
			maxSeatID = passport[i]
		}
	}
	//PART 1
	fmt.Println(maxSeatID)

	//PART 2
	sort.Ints(passport)
	for i:=0; i<len(passport); i++{
		if passport[i+1] - passport[i] == 2 {
			fmt.Println(passport[i]+1)
			break
		}
	}



}
