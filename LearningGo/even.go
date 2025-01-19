package main

import "fmt"

func even_number(arr []int) {

	for _, num := range arr{
		if num%2 == 0{
			fmt.Print(num)
		}
	}
}

func main() {
	var n int
	fmt.Print("Enter the size of array: ")
	fmt.Scan(&n)	

	arr := make([]int, n)

	fmt.Println("Enter the elements of array: ")
	for i:=0; i<n; i++{
		fmt.Scan(&arr[i])
	}
	even_number(arr)
}