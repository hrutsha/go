package main

import "fmt"

func main()  {
	var name string
	var age int

	fmt.Print("Enter your name : ")
	fmt.Scan(&name)

	fmt.Print("Enter your age : ")
	fmt.Scan(&age)

	fmt.Print("Your name :",name)
	fmt.Print("your age :",age)
}
