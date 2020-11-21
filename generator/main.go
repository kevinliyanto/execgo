package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	filecmd := exec.Command(execPath, newArgs...)

	// Use exec on the file
	fmt.Println(newArgs)

	// Read all from stdin. Will stop on EOF
	stdinData, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	// Pipe stdinData to exec
	// Read exec's stdout to a reader
	// Then take stdout to something called output

	// Open up WriteCloser and ReadCloser pipe to write and read from the exec'd command (Cmd)
	fWriteCloser, err := filecmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	fReadCloser, err := filecmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// Start the Cmd
	if err := filecmd.Start(); err != nil {
		log.Fatal(err)
	}

	// Pipe stdin to the Cmd, and read stdout of the Cmd to buf
	fWriteCloser.Write(stdinData)

	buf := new(bytes.Buffer)
	buf.ReadFrom(fReadCloser)

	if err = filecmd.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())

	// TODO
	// Make directory (recursively) based on path
	// If it is not an absolute path, use current path as base path
	// Else, use absolute path
	fmt.Println(dirPath)

	// TODO
	// Copy input and output data to file into the specified path

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
