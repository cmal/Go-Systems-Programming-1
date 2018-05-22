package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		os.Exit(1)
	}

	minusV := flag.Bool("v", false, "verbose mode")

	flag.Parse()
	flags := flag.Args()

	for _, file := range flags {
		err := os.Remove(file)
		if err != nil {
			fmt.Println(err)
			return
		}

		if *minusV && err == nil {
			fmt.Println("Removed file: ", file)
		}
	}
}
