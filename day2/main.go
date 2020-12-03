package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	/**
	  Part 1
	*/
	stop := false
	number := 0
	for scanner.Scan() {
		line := scanner.Text()
		x,_ := strconv.Atoi(line)
		file1, _ := os.Open("input.txt")
		scanner2 := bufio.NewScanner(file1)
		for scanner2.Scan(){
			line = scanner2.Text()
			y, _:= strconv.Atoi(line)
			if x + y == 2020 {
				number = x * y
				stop = true
				break
			}
			if stop {
				break
			}
		}
	}


	fmt.Print("Part 1: ")
	fmt.Println(number)

/**
Part 2
 */

	file, _ = os.Open("input.txt")
	scanner = bufio.NewScanner(file)
	/**
	  Part 1
	*/
	stop = false
	number = 0
	for scanner.Scan() {
		line := scanner.Text()
		x,_ := strconv.Atoi(line)
		file1, _ := os.Open("input.txt")
		scanner2 := bufio.NewScanner(file1)
		for scanner2.Scan() {
			line = scanner2.Text()
			y, _ := strconv.Atoi(line)
			file1, _ := os.Open("input.txt")
			scanner3 := bufio.NewScanner(file1)
				for scanner3.Scan(){
					line = scanner3.Text()
					z, _:= strconv.Atoi(line)
					if x + y + z == 2020 {
						number = x * y * z
						stop = true
						break
					}
					if stop {
						break
					}
				}
			if stop {
				break
			}
		}
	}
	fmt.Print("Part 2: ")
	fmt.Println(number)

}