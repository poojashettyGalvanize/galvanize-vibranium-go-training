package main

import (
	"fmt"
)

func main() {
	number1:=10
	number2:=15
	number3:=20
	sumOf2Numbersvalue:=sumOf2Numbers(number1,number2)
	fmt.Printf("Sum of numbers %v and %v is %v", number1,number2,sumOf2Numbersvalue)
	sumOf3Numbersvalue:=sumOf3Numbers(number1,number2,number3)
	fmt.Printf("\nSum of numbers %v, %v and %v is %v", number1,number2,number3,sumOf3Numbersvalue)
}

func sumOf2Numbers(value1,value2 int)int{
	sum:=value1+value2
	return sum
}

func sumOf3Numbers(value1,value2,value3 int)int{
	sum:=value1+value2+value3
	return sum
}