package main

import (
	"fmt"
	"os"
	"path/filepath"
)


func main(){

	// reading the cmd line args
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("please provide more arguments")
		return
	}

	//save it a variable called file
	file := arguments[1]
	// get content of the environment variable
	path := os.Getenv("PATH")
	// split the content 
	pathSplit := filepath.SplitList(path)
	// then iterate over all the directories of the path.
	for _, directory := range pathSplit{
		fullPath := filepath.Join(directory, file)
		// does the file exist.
		fileInfo,err := os.Stat(fullPath)

		if err == nil{
			mode := fileInfo.Mode()
			// is it a regular file.
			if mode.IsRegular(){
				// is it an executable file.
				if mode&0111 != 0{
					fmt.Println(fullPath)
					return
				}
			}
		}

	}



}