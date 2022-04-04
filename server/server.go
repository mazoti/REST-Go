package main

import (
	"fmt"
	"log"
	"os"

	"example.com/rest_server"
)

func main() {

	// Not enough arguments
	if len(os.Args) != 2 {
		fmt.Println(rest_server.USAGE_MSG)
		os.Exit(rest_server.USAGE_ERROR)
	}

	log.Println(rest_server.START_MSG)

	// Server only stops on error or when user press CTRL+C
	if err := rest_server.Start(os.Args[1]); err != nil {
		log.Println(err.Error())
		os.Exit(err.Code())
	}

	// No errors
	log.Println(rest_server.END_MSG)
}
