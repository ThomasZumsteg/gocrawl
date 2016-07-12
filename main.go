package main

import (
	"encoding/json"
	"flag"
    "io"
	"log"
    "os"
)

type Config struct {
    UserName string
    Password string
    Debug int
}


func main() {
	var backgroundPtr = flag.Bool("background", false, "Runs the command in the background")
	var databasePtr = flag.String("database", "./crawlnet", "Use this specified database")
	var verbosePtr = flag.Bool("verbose", false, "Run in verbose mode")
	var debugPtr = flag.Int("debug", 0, "Set the debug level (Default is 0)")
    var logFile = flag.String("log", "", "Set the log file location")
	flag.Parse()

    var logWriter io.Writer
    switch {
    case *logFile == "":
            logWriter = os.Stdout
    }


    logger := log.New(logWriter, "main: ", log.LstdFlags)

	logger.Print("Background in: ", *backgroundPtr)
	logger.Print("Database is: ", *databasePtr)
	logger.Print("Verbose is: ", *verbosePtr)
	logger.Print("Debug level is: ", *debugPtr)
	logger.Print("Roots are:")
	for _, root := range flag.Args() {
		logger.Printf("\t%s", root)
	}

    configFile, loggerErr := os.Open("config.json")
    if loggerErr != nil {
        logger.Print("Could not open config file")
        return
    }

    decoder := json.NewDecoder(configFile)
    config := Config{}
    decorderErr := decoder.Decode(&config)
    if decorderErr != nil {
        logger.Print("Config files appears to be incorrectly formatted")
        return
    }

    logger.Print("The username is: ", config.UserName)
    logger.Print("The password is: ", config.Password)
    logger.Print("The config debug level is: ", config.Debug)

	// If Database isn't a file, confirm overwrite
	//
}
