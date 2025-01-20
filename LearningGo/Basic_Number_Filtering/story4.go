package main

import(
	"fmt"
	"strings"
)

func oddPrime_number(arr []int) []int{
	result := []int{}
	for _,num := range arr{
		if num <= 1 {
			continue
		}
		if num == 2{
			result = append(result,num)}
		for i:=2; i<num; i++{
			if num%i == 0{
				break
			}
			if i == num-1{
				result = append(result,num)
			}
		}
	}
	odd_result := []int{}
	for _, num := range result{
		if num%2 != 0 {
			odd_result = append(odd_result,num)
		}
	}
	return odd_result
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
	Nums := oddPrime_number(arr)
	for i,num:= range Nums{
		if i==len(Nums)-1{
			fmt.Println(num)
		} else{
			fmt.Print(num,",")
		}
	}
}