package main

import (
	"strings"
	"fmt"
	"io"
	"os"
)


var out io.Writer = os.Stdout // modified during testing
var exit func(code int) = os.Exit

func main(){
	var input string
	fmt.Scan(&input)
	printNum(input)
}

func printNum(input string) {
	input = strings.ReplaceAll(input, ".", "")
	formatPrint := "%c%s\n"
	
	if strings.Contains(input, "-") {
		formatPrint = "-%c%s\n"
		input = strings.ReplaceAll(input, "-", "")
	}

	for i := 0; i < len(input); i++ {
		fmt.Fprintf(out, formatPrint, input[i], strings.Repeat("0", len(input)-1-i))
	}
}