package main

import "fmt"

func sum(a int,b int) int {
	return (a + b)
}

func main(){
	var num1 = 10
	var num2 = 5
	
	var sum = sum(num1,num2)
	fmt.Print(sum)	
}
