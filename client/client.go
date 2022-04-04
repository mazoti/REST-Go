package main

import (
	"fmt"
	"os"

	"example.com/rest_client"
)

func main() {

	// Not enough arguments
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, rest_client.USAGE_MSG)
		os.Exit(rest_client.USAGE_ERROR)
	}

	restClient, err := rest_client.Run(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(err.Code())
	}

	// No errors
	fmt.Println(restClient)
}
