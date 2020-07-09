package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
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
		fmt.Println("Usage: ", os.Args[0], "<directory-path> <master-executable>")
		return
	}
	file := os.Args[2]
	re, _ := regexp.Compile(`^(\w+)\.(\w+)$`)
	fi := re.FindAllStringSubmatch(file, -1)

	var execPath string = ""
	dirPath := os.Args[1]

	if len(fi[0]) != 3 {
		// Executable
		execPath = ""
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

	var newArgs []string

	if execPath == "" {
		execPath = file
		newArgs = os.Args[3:]
	} else {
		newArgs = append([]string{file}, os.Args[3:]...)
	}

	// filecmd := exec.Command(execPath, newArgs...)

	// Make directory based on path
	fmt.Println(dirPath)

	// Use exec on the file
	fmt.Println(newArgs)
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
	return path.Join(dir, config.Input), path.Join(dir, config.Output)
}
