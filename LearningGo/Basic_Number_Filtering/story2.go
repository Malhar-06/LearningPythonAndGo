package main

import("fmt"
"strings")

func odd_number(arr []int) []int{
	result := []int{}
	for num := range arr{
		if num%2 != 0 {
			result = append(result,num)
		}
	}
	return result
}


func main(){
	var input string
	fmt.Println("Enter elements of array (comma-seperated):")
	fmt.Scan(&input)

	arr := make([]int,0)

	numStrings := strings.Split(input,",")
	for _, numStr := range numStrings{
		var num int
		fmt.Sscanf(numStr, "%d", &num)
		arr = append(arr,num)
	}
	oddNums := odd_number(arr)
	for i,num:= range oddNums{
		if i==len(oddNums)-1{
			fmt.Println(num)
		} else{
			fmt.Print(num,",")
		}
	}
}