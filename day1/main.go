package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	/**
	Part 1
	 */
	validPasswords := 0
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ":")
		password := strings.TrimSpace(s[1])
		s = strings.Split(s[0], " ")
		char := s[1]
		s = strings.Split(s[0], "-")
		min, _:= strconv.Atoi(s[0])
		max, _ := strconv.Atoi(s[1])
		occurrences := strings.Count(password,char)
		if occurrences <= max && occurrences >= min {
			validPasswords++
		}
	}
	fmt.Print("Part 1: ")
	fmt.Println(validPasswords)
	file.Close()
	/**
	Part 2
	 */
	validPasswords = 0
	file, _ = os.Open("input.txt")
	scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ":")
		password := strings.TrimSpace(s[1])
		s = strings.Split(s[0], " ")
		char := s[1]
		s = strings.Split(s[0], "-")
		first, _ := strconv.Atoi(s[0])
		second, _ := strconv.Atoi(s[1])
		first--
		second--

		charPosition1 := string([]rune(password)[first])
		charPosition2 := string([]rune(password)[second])

		if charPosition1 != char && charPosition2 != char {
		}else if charPosition1 == char && charPosition2 == char {
		} else if charPosition1 == char || charPosition2 == char {
			validPasswords++
		}
	}
	fmt.Print("Part 2: ")
	fmt.Println(validPasswords)
}





