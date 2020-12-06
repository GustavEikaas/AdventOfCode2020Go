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

	for i:=0; i<len(answers); i++{
		answers[i] = re.ReplaceAllString(answers[i], "")
	}

	var numAnswers,uniqueChars int
	for i:=0; i<len(answers); i++{
		unique := make([]uint8,26)
		for j:=0; j<len(answers[i]); j++{
			char := answers[i][j]
			if!(existsInSlice(char,unique)){
				unique = append(unique, char)
				uniqueChars++
			}
		}
		numAnswers += uniqueChars
		uniqueChars = 0
		unique = nil
	}
	fmt.Println(numAnswers)

	/*
	Part 2
	 */
	input,_ = ioutil.ReadFile("input.txt")
	answers = strings.Split(string(input),"\n\n")

	var part2 int
	for i:=0; i<len(answers); i++{
		var repeat int
		indvAnswer := strings.Split(answers[i],"\n")

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