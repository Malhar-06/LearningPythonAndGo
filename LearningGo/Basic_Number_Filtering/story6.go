package main

import(
	"fmt"
	"strings"
)

func oddMultipleOf3grth10_number(arr []int) []int{
	multiple3_result := []int{}

	for _, num := range arr{
		if num%2!=0 && num%3 == 0 && num > 10 {
			multiple3_result = append(multiple3_result,num)
		}
	}
	return multiple3_result
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
	Nums := oddMultipleOf3grth10_number(arr)
	for i,num:= range Nums{
		if i==len(Nums)-1{
			fmt.Println(num)
		} else{
			fmt.Print(num,",")
		}
	}
}