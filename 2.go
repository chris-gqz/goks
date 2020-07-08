package main

import "fmt"

// 代码规范
func main() {
	//var numsStr string
	numsStr := "631758924"
	//fmt.Scanln(&numsStr)
	var decodeNums []byte
	var tmp []byte
	count := 0
	for byteNums := []byte(numsStr); len(byteNums) != 0; byteNums = tmp {
		tmp = nil
		for k, num := range byteNums {
			k = count + k
			if k%2 == 1 {
				tmp = append(tmp, num)
				continue
			}
			decodeNums = append(decodeNums, num)
		}
		count = len(byteNums) + count
	}
	fmt.Println(string((decodeNums)))
}
