package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main(){
	/*
	Part 1
	 */
	input,_ := ioutil.ReadFile("input.txt")
	answers := strings.Split(string(input),"\n\n")
	re := regexp.MustCompile(`\r?\n`)
	part1Answers := make([]string,len(answers))
	for i:=0; i<len(answers); i++{
		part1Answers[i] = re.ReplaceAllString(answers[i], "")
	}
	var part1,uniqueChars int
	for i:=0; i<len(part1Answers); i++{
		unique := make([]uint8,26)
		for j:=0; j<len(part1Answers[i]); j++{
			char := part1Answers[i][j]
			if!(existsInSlice(char,unique)){
				unique = append(unique, char)
				uniqueChars++
			}
		}
		part1 += uniqueChars
		uniqueChars = 0
		unique = nil
	}
	fmt.Println(part1)

	/*
	Part 2
	 */
	var part2,repeat int
	for i:=0; i<len(answers); i++{
		indvAnswer := strings.Split(answers[i],"\n")
		lastLine := len(indvAnswer)-1
		if indvAnswer[lastLine] == "" {
			indvAnswer = indvAnswer[0:lastLine-1]
		}
		for j:=0; j<len(indvAnswer[0]); j++{
			char := string([]rune(indvAnswer[0])[j])
			for k:=0; k<len(indvAnswer); k++{
				if strings.Contains(indvAnswer[k],char){
					repeat++
				}
			}
			if repeat == len(indvAnswer) {
				part2++
			}
			repeat = 0
		}
	}
	fmt.Println(part2)
}

func existsInSlice(char uint8, unique []uint8)bool {
	for k := 0; k < len(unique); k++ {
		if char == unique[k] {
			return true
		}
	}
	return false
}