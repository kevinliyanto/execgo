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

## Generate test executable

1. `chmod u+x script.pl`
2. `./script.pl <command>`
   Optional commands are:

- (no command: generates `./test`)
- all (generates `./generator` alongside `./test`)
- clean (clean all executable files)
- generator (only generates `./generator`)

## Config.json file

This config.json file indicates where the test input and expected output file of a directory that will be specified on `./test`.

### Todos

- Implement diff on `executor/diff.go`
