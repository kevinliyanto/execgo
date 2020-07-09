package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// Configuration for expected input file
type Configuration struct {
	Input  string
	Output string
}

var executablesMap map[string]string

func init() {
	executablesMap = map[string]string{
		"py": "python",
		"js": "node",
		"pl": "perl",
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ", os.Args[0], "<directory> <code>")
		return
	}
	file := os.Args[2]
	re, _ := regexp.Compile(`^(\w+)\.(\w+)$`)
	fi := re.FindAllStringSubmatch(file, -1)

	var execPath string = ""

	if len(fi[0]) != 3 {
		// File is an executable
		// Use empty execPath
	} else {
		// File to interpret
		executable, exist := executablesMap[fi[0][2]]

		if exist {
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
		} else {
			fmt.Println("Filetype is not supported")
			return
		}
	}

	// Get input/output file
	config, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		log.Fatal(err)
		os.Exit(1)
	}

	input, output := GeneratePath(config, os.Args[1])

	var newArgs []string

	if execPath == "" {
		execPath = file
		newArgs = os.Args[3:]
	} else {
		newArgs = append([]string{file}, os.Args[3:]...)
	}

	filecmd := exec.Command(execPath, newArgs...)

	// Open input file and stream it to the exec's stdin
	inputFile, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer inputFile.Close()

	filecmd.Stdin = inputFile
	filecmd.Stderr = os.Stderr

	originalOutput, err := filecmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if err := filecmd.Start(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(originalOutput)

	if err = filecmd.Wait(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	programOutput := buf.String()

	outputFile, err := ioutil.ReadFile(output)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	outputText := string(outputFile)

	if equal, diff := Diff(outputText, programOutput); equal {
		fmt.Println("Output is right")
	} else {
		fmt.Printf("Output is wrong. Difference: %s\n", diff)
	}
}

// ReadConfig get read configuration input and output filename
func ReadConfig() (*Configuration, error) {
	file, _ := os.Open("config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Configuration{}
	err := decoder.Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// GeneratePath generates path from the given directory with the config specified
func GeneratePath(config *Configuration, dir string) (string, string) {
	return filepath.Join(dir, config.Input), filepath.Join(dir, config.Output)
}
