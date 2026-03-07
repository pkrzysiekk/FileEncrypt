package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if args := os.Args; len(args) < 2 {
		log.Fatal("Not enough arguments, provide filename and password")
	}
	filename, password := os.Args[1], os.Args[2]
	fmt.Println(filename, password)

}
