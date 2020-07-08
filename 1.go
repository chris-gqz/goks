package main

import (
	"fmt"
	"strconv"
	"strings"
)

var priorityMap = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

func main() {
	var calcString string
	fmt.Scanln(&calcString)
	//calcString := "[1234]=[12]+[34]*{50},[12]=[1]+[2]/{2};[1]=10,[2]=20,[34]=50;[1234]"
	//calcString := "[10001]=[1001]+[1002]-[1001];[1001]=10,[1002]=20;[10001]"
	//calcString := "[10004]=[1001]/{10}+[1002]*{10};[1001]=10,[1002]=20;[10004]"
	res, all := handleInputStr(calcString)
	fmt.Println(calc(res, all))
}

func handleInputStr(inputStr string) (string, map[string]interface{}) {
	inputStr = strings.Replace(inputStr, "[", "", -1)
	inputStr = strings.Replace(inputStr, "]", "", -1)
	inputList := strings.Split(inputStr, ";")
	res := inputList[2]
	var all = make(map[string]interface{})
	inputListItem1 := strings.Split(inputList[1], ",")
	for _, item := range inputListItem1 {
		leftRight := strings.Split(item, "=")
		leftRightInt, _ := strconv.Atoi(leftRight[1])
		all[leftRight[0]] = leftRightInt
	}

	inputListItem0 := strings.Split(inputList[0], ",")
	for _, item := range inputListItem0 {
		leftRight := strings.Split(item, "=")
		all[leftRight[0]] = leftRight[1]
	}
	return res, all
}

func calc(res string, all map[string]interface{}) int {
	v, ok := all[res].(int)
	if ok {
		return v
	}
	numbers := []int{}
	operators := []string{}
	number := ""
	valueStr, ok := all[res].(string)
	if !ok {
		return -1
	}
	for _, ch := range valueStr {
		//0==>48,9==>57
		if 48 <= ch && ch <= 57 {
			number += string(ch)
		} else if string(ch) == "}" {
			tmpNumber, _ := strconv.Atoi(number)
			numbers = append(numbers, tmpNumber)
			number = ""
		} else {
			if _, ok := priorityMap[string(ch)]; ok {
				if number != "" {
					numbers = append(numbers, calc(number, all))
					number = ""
				}
				if len(operators) > 0 && priorityMap[string(ch)] <= priorityMap[operators[len(operators)-1]] {
					tmpNumber := calcTwoNumber(numbers, operators)
					numbers = numbers[:len(numbers)-2]
					operators = operators[:len(operators)-1]
					numbers = append(numbers, tmpNumber)
				}
				operators = append(operators, string(ch))
			}
		}
	}

	if len(number) > 0 {
		numbers = append(numbers, calc(number, all))
	}

	for ; len(operators) > 0; {
		tmp := calcTwoNumber(numbers, operators)
		numbers = numbers[:len(numbers)-2]
		operators = operators[:len(operators)-1]
		numbers = append(numbers, tmp)
	}
	all[res] = numbers[len(numbers)-1]

	return numbers[len(numbers)-1]
}

func calcTwoNumber(nums []int, opers []string) int {
	right := nums[len(nums)-1]
	left := nums[len(nums)-2]
	oper := opers[len(opers)-1]

	if oper == "-" {
		return left - right
	} else if oper == "*" {
		return right * left
	} else if oper == "/" {
		return left / right
	}
	return right + left
}
