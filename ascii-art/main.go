package main

import (
	"fmt"
	"os"
	"strings"

	asciiart "asciiart/src"
)

func main() {

	// The function checks the length of the command-line arguments passed to
	// the program using len(os.Args[1:]).

	// If there is only one argument (excluding the program name itself),
	// it assumes that the argument is the input string to be converted to ASCII art.
	// It calls the AsciiPrint function from the asciiart package,
	// passing the input string with any occurrences of the literal \n replaced with
	// actual newline characters (\n).

	if len(os.Args[1:]) == 1 {
		asciiart.AsciiPrint(strings.ReplaceAll(os.Args[1], `\n`, "\n"), "standard")

		// If there are two arguments, it assumes that the first argument is
		// the input string and the second argument is the font to be used for
		// the ASCII art. It calls the AsciiPrint function from the asciiart package,
		// passing the input string with any occurrences of the literal \n replaced with
		// actual newline characters (\n), and the font provided as the second argument.

	} else if len(os.Args[1:]) == 2 {
		asciiart.AsciiPrint(strings.ReplaceAll(os.Args[1], `\n`, "\n"), os.Args[2])

		// If there are neither one nor two arguments, it prints a usage message that
		// explains the expected command-line arguments for the program.

	} else {
		fmt.Println("Usage: go run main.go \"[input]\"")
	}
}
