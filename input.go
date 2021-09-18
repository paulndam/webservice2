package main

import (
	"fmt"
)
 func main(){
	 // get user input.
	 fmt.Printf("please enter your name: ")
	 var name string 
	 fmt.Scanln(&name)
	 fmt.Println("your name is :---->", name)
 }