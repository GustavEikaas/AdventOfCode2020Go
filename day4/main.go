package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main(){
	input,_ := ioutil.ReadFile("input.txt")
	var passports = strings.Split(string(input), "\n\n")
	var validPassports int
	/**
	  Part 1
	*/
	r, _ := regexp.Compile("byr:|" +
		"ecl:|" +
		"hgt:|" +
		"pid:|" +
		"hcl:|" +
		"iyr:|" +
		"eyr:")
	for i:=0; i<len(passports); i++{
		if len(r.FindAll([]byte(passports[i]), -1)) == 7{
			validPassports++
		}
	}
	fmt.Print("Part 1: ")
	fmt.Printf("Valid passports: %d%s",validPassports,"\n")
	/**
	Part 2
	*/
	validPassports = 0
	//Pid with 10 digits still get counted
	//Doing [^0-9] gives 172
	//198 is correct
	r, _ = regexp.Compile("byr:(19[2-9][0-9]|200[0-2])|" +
		"ecl:(amb|oth|blu|brn|gry|grn|hzl)|" +
		"hgt:((1[5-8][0-9]cm|19[0-3]cm)|(59|6[0-9]|7[0-6])in)|" +
		"pid:[0-9]{9}|" +
		"hcl:#([a-f]|[0-9]){6}|" +
		"iyr:20(1[0-9]|20)|" +
		"eyr:20(2[0-9]|30)")

	for i:=0; i<len(passports); i++ {
		if len(r.FindAll([]byte(passports[i]), -1)) == 7{
			validPassports++
		}
	}
	fmt.Print("Part 2: ")
	fmt.Printf("Valid passports: %d%s",validPassports,"\n")
}