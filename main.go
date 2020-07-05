package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
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
	if len(os.Args) < 3 {
		fmt.Println("Usage: ", os.Args[0], "<directory> <code>")
		return
	}
	file := os.Args[2]
	re, _ := regexp.Compile(`^(\w+)\.(\w+)$`)
	fi := re.FindAllStringSubmatch(file, -1)

	if len(fi[0]) != 3 {
		// Executable
		fmt.Println(fi[0][0])
	} else {
		// File to interpret
		var execPath string

		if executable, exist := executablesMap[fi[0][2]]; exist {
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

		newArgs := append([]string{file}, os.Args[3:]...)
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
}
