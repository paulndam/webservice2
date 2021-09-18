package main

import (
	"log"
	"os"
)


func main(){
	if len(os.Args) != 1 {
		log.Fatal("Fatal : Hello")
	}

	log.Panic("Panic : hello rico")
}