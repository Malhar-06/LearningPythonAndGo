package main

import(
	"fmt"
	"strings"
)

func CustomOperations_number(arr []int, operationArr []int) []int{
	result := []int{}
	for _, operation := range operationArr {
	switch operation{

	case 1:
		for _, num := range arr{
			if num%2 == 0 {
				result = append(result,num)
			}
		}
	case 2:
		for _, num := range arr{
			if num%2 != 0 {
				result = append(result,num)
			}
		}
	case 3:
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
	case 4:
		for _, num := range arr{
			if num%5 == 0 {
				result = append(result,num)
			}
		}
	case 5:
		for _, num := range arr{
			if num%3 == 0 {
				result = append(result,num)
			}
		}
	case 6:
		for _, num := range arr{
			if num > 10 {
				result = append(result,num)
			}
		}
	}}
	return result
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
	
	fmt.Println("Select Which Operations you want to perform:")
	fmt.Println("1. Even")
	fmt.Println("2. Odd")
	fmt.Println("3. Prime")
	fmt.Println("4. Multiple of 5")
	fmt.Println("5. Multiple of 3")
	fmt.Println("6. Greater than 10")
	fmt.Println("Select the operation by entering the numbers (eg.1,3,5):")
	var operation string
	fmt.Scan(&operation)
	input = strings.ReplaceAll(input, " ", "") // Remove spaces
	operationStrings := strings.Split(operation, ",")
	operationArr := make([]int,0)
	for _,operationStr:= range operationStrings{
		var num int
		fmt.Sscanf(operationStr, "%d", &num)
		operationArr = append(operationArr,num)
	}

	Nums := CustomOperations_number(arr,operationArr)
	for i,num:= range Nums{
		if i==len(Nums)-1{
			fmt.Println(num)
		} else{
			fmt.Print(num,",")
		}
	}
}