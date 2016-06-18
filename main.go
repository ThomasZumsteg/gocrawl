package main

import (
	// "./gocrawl"
	// "encoding/json"
	"flag"
	"fmt"
	// "log"
)

func main() {
	var backgroundPtr = flag.Bool("background", false, "Runs the command in the background")
	var databasePtr = flag.String("database", "./crawlnet", "Use this specified database")
	var verbosePtr = flag.Bool("verbose", false, "Run in verbose mode")
	var debugPtr = flag.Int("debug", 0, "Set the debug level (Default is 0)")

	flag.Parse()

	fmt.Println("Background in: ", *backgroundPtr)
	fmt.Println("Database is: ", *databasePtr)
	fmt.Println("Verbose is: ", *verbosePtr)
	fmt.Println("Debug level is: ", *debugPtr)
	fmt.Println("Roots are:")
	for _, root := range flag.Args() {
		fmt.Printf("\t%s\n", root)
	}

	// If Database isn't a file, confirm overwrite
	//
}
