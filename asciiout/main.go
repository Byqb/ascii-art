package main

import (
	out "out/Functions"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) == 2 {
		out.AsciiPrint(strings.ReplaceAll(os.Args[1], `\n`, "\n"), os.Args[2])
	} else if len(os.Args[1:]) == 3 && strings.HasPrefix(os.Args[3], "--output=") {
		out.AsciiPrint(strings.ReplaceAll(os.Args[1], `\n`, "\n"), os.Args[2], os.Args[3])
	} else {
		fmt.Println(
			"Usage: go run . [STRING] [BANNER] [OPTION]\n\nEX: go run . \"something\" standard --output=<fileName.txt>")
	}
}
