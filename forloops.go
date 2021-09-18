package main

import (
	"fmt"
)

func main()  {
// traditional for loop.
fmt.Println("traditional for loop")

for i := 0; i < 10; i++{
	fmt.Println(i * i, " ")
}	

//idiomatic Go
fmt.Println("idiomatic loop")

i := 0
for ok := true; ok;ok = (i != 10){
	fmt.Println(i* i," ")
	i++
}
fmt.Println()

// for loop used as while loop.
fmt.Println("while loop")
j := 0
for{
	if i == 10{
		break
	}
	fmt.Println(j * j, " ")
	j++
}
fmt.Println()



// using range in a slice array.
aSlice := []int {-1,2,3,-6,-7,-8}
for index,value := range aSlice{
	fmt.Println("index--->",index, "value:--->",value)
}


}