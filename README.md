# Execgo

A tool to execute an executable/interpreter, either comparing it to a specific input/output or generates input/output

## Executor

Compares output of a program to a given output file with the specified input file.

## Generator

Based on your executables / interpreted file, generator will generate input (from stdin) and output (from stdout) on the specified directory.

## How to init

Make sure that these are available (check by using `which`):

- go
- make
- perl

## Supported interpreters:

- python
- node
- perl

## Generate executable

1. `make`
2. `./output <test|generate> <file> [ --options="values" ]`

| Required | Default         | Option (angled bracket not required)         | Description                                                                   |
| -------- | --------------- | -------------------------------------------- | ----------------------------------------------------------------------------- |
| Yes      | N/A             | First parameter (test or generate            | Generate input/output file or compare runnable to input/output file reference |
| Yes      | N/A             | Second parameter (file)                      | File that you're going to run. Can be interpreted file or direct executable   |
| No       | "" (empty)      | `-a="<arguments>"` or `--args="<arguments>"` | Arguments supplied to the runnable                                            |
| No       | "file_solution" | `-p="<path>"` or `--path="<path>"`           | Custom path which points to directory which stores input and output file      |
| No       | "input"         | `-i="<filename>"` or `--input="<filename>"`  | Input file reference name                                                     |
| No       | "output"        | `-o="<filename>"` or `--output="<filename>"` | Output file reference name                                                    |
