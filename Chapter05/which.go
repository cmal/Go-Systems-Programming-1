package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	minusA := flag.Bool("a", false, "a")
	minusS := flag.Bool("s", false, "s")

	flag.Parse()
	flags := flag.Args()
	if len(flags) == 0 {
		fmt.Println("Please provide an argument!")
		os.Exit(1)
	}

	allFound := true
	for _, file := range flags {
		fountIt := false
		path := os.Getenv("PATH")
		pathSlice := strings.Split(path, ":")
		for _, directory := range pathSlice {
			fullPath := directory + "/" + file
			// Does it exist?
			fileInfo, err := os.Stat(fullPath)
			if err == nil {
				mode := fileInfo.Mode()
				// Is it a regular file?
				if mode.IsRegular() {
					// Is it executable?
					// fmt.Println(mode, mode&0111, mode&0777)
					if mode&0111 != 0 {
						fountIt = true
						// if the -s flag is set
						if *minusS == true {
							// os.Exit(0)
							continue
						}
						// if the -a flag is set
						if *minusA == true {
							fmt.Println(fullPath)
						} else {
							fmt.Println(fullPath)
							// os.Exit(0)
						}
					}
				}
			}
		}
		if fountIt == false {
			fmt.Println(file, "not found")
			allFound = false
		}

	}

	if allFound == false {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
