package main

import (
	"bytes"
	"fmt"
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
	if len(os.Args) < 2 {
		fmt.Println("Usage: ", os.Args[0], "<FILE>")
		return
	}
	file := os.Args[1]
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
			fmt.Println(execPath)
		} else {
			fmt.Println("Filetype is not supported")
			return
		}

		newArgs := append([]string{file}, os.Args[2:]...)
		filecmd := exec.Command(execPath, newArgs...)
		// Pipe in from input file
		filecmd.Stdin = os.Stdin
		//
		filecmd.Stdout = os.Stdout
		filecmd.Stderr = os.Stderr

		output, err := filecmd.Output()
		if err != nil {
			log.Fatal(err)
		}

		if err := filecmd.Run(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Output: %s", output)
	}
}
