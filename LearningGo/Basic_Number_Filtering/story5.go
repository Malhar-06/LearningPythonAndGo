package main

import(
	"fmt"
	"strings"
)

func evenMultipleOf5_number(arr []int) []int{
	result := []int{}
	multiple5_result := []int{}

	for _, num := range arr{
		if num%2 == 0 {
			result = append(result,num)
		}
	}
	for _, num := range result{
		if num%5 == 0 {
			multiple5_result = append(multiple5_result,num)
		}
	}
	return multiple5_result
}

func main(){
	var input string
	fmt.Println("Enter the numbers separated by commas:")
	fmt.Scan(&input)

	numStrings := strings.Split(input, ",")
	arr := make([]int,0)
	for _,numStr:= range numStrings{
		var num int
		fmt.Sscanf(numStr, "%d", &num)
		arr = append(arr,num)
	}
	Nums := evenMultipleOf5_number(arr)
	for i,num:= range Nums{
		if i==len(Nums)-1{
			fmt.Println(num)
		} else{
			fmt.Print(num,",")
		}
	}
}