// Package Zconvert :
// File:  Zconvert.go
// Date:  2021/1/19 22:03
// Desc:
package Zconvert

func Zconvert(str string, numRows int) string {
	runes := make([][]rune, numRows)
	cycle := 2*numRows - 2 // 通过数学找出规律
	for i, char := range str {
		index := i % cycle
		if index < numRows {
			runes[index] = append(runes[index], char)
		} else {
			runes[cycle-index] = append(runes[cycle-index], char)
		}
	}
	result := ""
	for _, str := range runes {
		result += string(str)
	}
	return result
}
