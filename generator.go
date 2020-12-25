package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Generate generates reference input and output for the runnable, storing in dirPath
func Generate(config *Configuration) {
	filecmd := config.exec
	dirPath := config.optPath

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

	stdoutDataBuf := new(bytes.Buffer)
	stdoutDataBuf.ReadFrom(fReadCloser)

	if err = filecmd.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(stdoutDataBuf.String())

	// Make directory (recursively) based on path
	if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	inputPath := filepath.Join(dirPath, config.inputFile)
	outputPath := filepath.Join(dirPath, config.outputFile)

	// TODO
	// Copy input and output data to file into the specified path

	// Write input
	if err = ioutil.WriteFile(inputPath, stdinData, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Write output
	if err = ioutil.WriteFile(outputPath, stdoutDataBuf.Bytes(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
