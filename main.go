package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

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

// Configuration ...
type Configuration struct {
	exec       *exec.Cmd
	optPath    string
	inputFile  string
	outputFile string
}

func main() {
	// Parse first and second argument to the the argument

	// First as mode (test or generate)
	mode := -1

	if len(os.Args) < 3 {
		fmt.Printf("usage: %s <test|generate> <file> [ --options=\"values\" ]\n", os.Args[0])
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

	// Check if file can be run
	re, _ := regexp.Compile(`(\w+)\.(\w+)$`)
	fi := re.FindAllStringSubmatch(file, -1)

	// execPath is needed to get the command that needs to be run
	var execPath string

	if len(fi) == 0 {
		// File is the executable
		execPath = file
	} else {
		// Get runtime if exists
		executable, exist := executablesMap[fi[0][2]]
		// Also, set file to the filename without the extension
		file = fi[0][1]
		if !exist {
			// File is not supported
			fmt.Println("Filetype is not supported")
			os.Exit(1)
		} else {
			// Find the executable
			cmd := exec.Command("which", executable)
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				log.Fatal(err)
			}
			if err := cmd.Start(); err != nil {
				log.Fatal(err)
			}
			buf := new(bytes.Buffer)
			buf.ReadFrom(stdout)
			execPath = strings.TrimSuffix(buf.String(), "\n")
		}
	}

	// --args flag to set the passed arguments, in a quoted string
	var passedArgs string
	getopt.FlagLong(&passedArgs, "args", 'a', "Arguments").SetOptional()
	args := strings.Split(passedArgs, "")

	cmd := exec.Command(execPath, args...)

	// --path flag to set custom filepath for executor/generator directory
	var optPath string
	getopt.FlagLong(&optPath, "path", 'p', "Path").SetOptional()

	// --input flag to set custom filepath for executor/generator input file
	var inputFile string
	getopt.FlagLong(&inputFile, "input", 'i', "Input").SetOptional()

	// --output flag to set custom filepath for executor/generator output file
	var outputFile string
	getopt.FlagLong(&outputFile, "outputFile", 'o', "Output").SetOptional()

	// Parse opts
	var opts = getopt.CommandLine
	opts.Parse(os.Args[2:])

	// If optPath is not set, set to default
	// (name of file without extension as directory, appended with `_solution`)
	if len(optPath) == 0 {
		optPath = file + "_solution"
	}

	// If input or putput is not set, set to default
	if len(inputFile) == 0 {
		inputFile = "input"
	}

	if len(outputFile) == 0 {
		outputFile = "output"
	}

	fmt.Println(optPath)
	fmt.Println(execPath)
	fmt.Println(cmd)

	config := &Configuration{
		exec:       cmd,
		optPath:    optPath,
		inputFile:  inputFile,
		outputFile: outputFile,
	}

	switch mode {
	case 0:
		// Call test with the params
		test(config)
	case 1:
		// Call generate with the params
	}
}

func test(config *Configuration) {
	// Dummy for test
	return
}

func generate(config *Configuration) {
	// Dummy for configuration
	return
}
