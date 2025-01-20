package main

import (
	"fmt"
	"strings"
)

func even_number(arr []int) []int {
	result := []int{} 
	for _, num := range arr{
		if num%2 == 0{
			result = append(result, num)
		}
	}
	return result
// 	for _, num := range result{
// 		fmt.Print(num, ',')
// 	}
}

func main() {
	var input string
	fmt.Print("Enter array elements (comma-seperated): ")
	fmt.Scanln(&input)
	
	// input = strings.ReplaceAll(input, " ", "") // Remove spaces
    numStrings := strings.Split(input, ",")

	arr := make([]int, 0)

	for _, numStr := range numStrings{
		var num int
		fmt.Sscanf(numStr, "%d", &num)
		arr = append(arr, num)
	}
	evennums := even_number(arr)
	for i, num := range evennums{
		if i==len(evennums)-1{
			fmt.Print(num)
		} else {
			fmt.Print(num, ",")
		}
	}
	fmt.Println()
}