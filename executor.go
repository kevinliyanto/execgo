package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Test compares executed file input/output to reference
func Test(config *Configuration) {
	filecmd := config.exec
	dirPath := config.optPath

	inputPath := filepath.Join(dirPath, config.inputFile)
	outputPath := filepath.Join(dirPath, config.outputFile)

	// Open input file and stream it to the exec's stdin
	inputFile, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	filecmd.Stdin = inputFile
	filecmd.Stderr = os.Stderr

	originalOutput, err := filecmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := filecmd.Start(); err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(originalOutput)

	if err = filecmd.Wait(); err != nil {
		log.Fatal(err)
	}

	programOutput := buf.String()

	outputFile, err := ioutil.ReadFile(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	outputText := string(outputFile)

	if equal, diff := Diff(outputText, programOutput); equal {
		fmt.Println("Output is right")
	} else {
		fmt.Printf("Output is wrong. Difference: %s\n", diff)
	}
}
