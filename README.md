# Execgo

A tool to execute an executable/interpreter, either comparing it to a specific input/output or generates input/output

**Test**: compares output of a program to a given output file with the specified input file.

**Generate**: based on your executables / interpreted file, generator will generate input (from stdin) and output (from stdout) on the specified directory.

## Supported interpreters:

- python
- node
- perl

## Generate executable

Make sure that these are available (check by using `which`):

- go
- make

```sh
go mod vendor
make
```

## Use executable

```sh
./output <test|generate> <file> [ --options="values" ]
```

Reference table for executable parameters and options:

| Required | Default         | Option (quotes are not required unless whitespace exists) | Description                                                                   |
| -------- | --------------- | --------------------------------------------------------- | ----------------------------------------------------------------------------- |
| Yes      | N/A             | First parameter (test or generate)                        | Generate input/output file or compare runnable to input/output file reference |
| Yes      | N/A             | Second parameter (file)                                   | File that you're going to run. Can be interpreted file or direct executable   |
| No       | "" (empty)      | `-a="<arguments>"` or `--args="<arguments>"`              | Arguments supplied to the runnable                                            |
| No       | "file_solution" | `-p="<path>"` or `--path="<path>"`                        | Custom path which points to directory which stores input and output file      |
| No       | "input"         | `-i="<filename>"` or `--input="<filename>"`               | Input file reference name                                                     |
| No       | "output"        | `-o="<filename>"` or `--output="<filename>"`              | Output file reference name                                                    |

## Delete executable

```sh
make clean
```
