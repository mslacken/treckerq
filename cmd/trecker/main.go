package main

import (
	"fmt"
	"os"

	"github.com/mslacken/treckerq/internal/app/trecker"
)

func main() {

	root := trecker.GetRootCommand()

	err := root.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(255)
	}
}
