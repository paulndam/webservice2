package main

import (
	"fmt"
	"os"
	"strconv"
)


func main()  {

	if len(os.Args) != 2 {
		fmt.Println("please input a command line argument")
		return
	}

	argument := os.Args[1]

//with expressions after switch.
switch argument{
case "0":
	fmt.Println("zero")
case "1":
	fmt.Println("one")
case "2":
	fmt.Println("two")
case "3", "4", "5":
	fmt.Println("three or four or five")
	fallthrough
default:
	fmt.Println("Value : --->", argument)

}

// convert input that is been recieved as string to integer.
//if error occurs then we throw error message.

value,err := strconv.Atoi(argument)
if err	!= nil{
	fmt.Println("can't convert to integer", argument)
	return
}

switch  {
case value == 0:
	fmt.Println("zero")
case value > 0:
	fmt.Println("positve integer")
case value < 0:
	fmt.Println("negative integer")
default:
	fmt.Println("i dont think this should be happening:",value)

	
}
	
}

