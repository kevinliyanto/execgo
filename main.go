package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pborman/getopt/v2"
)

var executablesMap map[string]string

func init() {
	executablesMap = map[string]string{
		"py": "python",
		"js": "node",
		"pl": "perl",
	}
}

func main() {
	// Parse first and second argument to the the argument

	// First as mode (test or generate)
	mode := -1

	if len(os.Args) < 3 {
		fmt.Printf("usage: %s <test|generate> <file>\n", os.Args[0])
		os.Exit(1)
	}

	switch os.Args[1] {
	case "test", "t":
		mode = 0
	case "generate", "g":
		mode = 1
	default:
		fmt.Println("expected 'test' or 'generate' subcommands")
		os.Exit(1)
	}

	// Second as file
	file := os.Args[2]

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Use filePath as something that gets executed
	filePath := filepath.Join(cwd, file)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println(err)
		os.Exit(1)
	}

	// --path flag to set custom filepath for executor/generator directory
	var optPath string
	getopt.FlagLong(&optPath, "path", 'p', "Path").SetOptional()

	// --args flag to set the passed arguments, in a quoted string
	var passedArgs string
	getopt.FlagLong(&passedArgs, "args", 'a', "Arguments").SetOptional()

	// --input flag to set custom filepath for executor/generator input file
	var inputFile string
	getopt.FlagLong(&inputFile, "input", 'i', "Input").SetOptional()

	// --output flag to set custom filepath for executor/generator output file
	var outputFile string
	getopt.FlagLong(&outputFile, "outputFile", 'o', "Output").SetOptional()

	var opts = getopt.CommandLine

	fmt.Println(os.Args[2])

	// Parse opts
	opts.Parse(os.Args[2:])

	// If optPath is not set, set to default (name of file as directory)
	if len(optPath) == 0 {

	}

	// If input or putput is not set, set to default
	if len(inputFile) == 0 {
		inputFile = "input"
	}

	if len(outputFile) == 0 {
		outputFile = "output"
	}

	fmt.Println(len(optPath))
	fmt.Println(optPath)
	fmt.Println(passedArgs)

	switch mode {
	case 0:
		// Call test with the params
	case 1:
		// Call generate with the params
	}
}
