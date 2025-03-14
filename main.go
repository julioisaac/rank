package main

import (
	"log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Missing arg")
	}

}
