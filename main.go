package main 

import (
	 
	"fmt"
	"flag"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
 	args := flag.Args()

	// requesting data if none is specified as a flag
	if len(args) < 1 {
	fmt.Println("Please enter data for conversion")  
	os.Exit(1)  
	}

	data := []byte(args[0])
	hash256 := crypto.newHash256(data)

}