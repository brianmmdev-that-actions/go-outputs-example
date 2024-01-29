package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		log.Printf("Hello %v!!", args[1])
		output := os.Getenv("GITHUB_OUTPUT")
		log.Println(output)
		os.Setenv("GITHUB_OUTPUT", fmt.Sprintf("\noutname=%v", args[1]))
		return
	}
	log.Println("Hello world!!")
}
