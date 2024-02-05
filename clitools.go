package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

func main() {

	//Bool Flag Created
	verbose := flag.Bool("v", false, "Verbose Mode Activated")

	//String Flag created
	name := flag.String("name", "World", "Name to be printed")

	n := flag.Int("i", 1, "Number of times")

	duration := flag.Duration("t", 5*time.Second, "duration....")
	flag.Parse()

	if *verbose {
		*name = strings.ToUpper(*name)
	}

	fmt.Printf("Hello, %s!\n", *name)
	for i := 0; i < *n; i++ {
		fmt.Print(*name)
	}

}
