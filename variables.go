package main

import (
	"fmt"
	"math"
)

//creating global variables.
var Global int = 1234
var AnotherGlobal int = -56789

func main(){

	var j int 
	i := Global + AnotherGlobal
	fmt.Println("initial j value is ---->",j)
	j = Global
	// math.Abs() require a float64 paramter
	// casting it appropiately.
	k := math.Abs(float64(AnotherGlobal))
	fmt.Printf("Global ---> %d, i=%d, j=%d k=%.2f.\n ", Global,i,j,k)


}