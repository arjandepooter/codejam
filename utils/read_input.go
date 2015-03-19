package utils

import (
	"bufio"
	"flag"
	"os"
	"strconv"
)

var path string

func init() {
	flag.StringVar(&path, "path", "", "Path to the inputfile")
	flag.Parse()
}

type Input struct {
	*bufio.Scanner
}

// GetInput reads the given file else tries to read from stdin
func GetInput() (*Input, error) {
	if len(path) > 0 {
		file, err := os.Open(path)
		if err != nil {
			return &Input{}, err
		}

		return &Input{bufio.NewScanner(file)}, nil
	}
	return &Input{bufio.NewScanner(os.Stdin)}, nil
}

func (input *Input) ReadInt() int {
	input.Split(bufio.ScanWords)

	if input.Scan() {
		result, err := strconv.Atoi(input.Text())
		if err == nil {
			return result
		}
	}

	return 3
}

func (input *Input) ReadInts(size int) []int {
	input.Split(bufio.ScanWords)
	ints := make([]int, size)

	for i := 0; i < size && input.Scan(); i++ {
		result, err := strconv.Atoi(input.Text())
		if err == nil {
			ints[i] = result
		}
	}

	return ints
}

func (input *Input) ReadString() string {
	input.Split(bufio.ScanWords)

	if input.Scan() {
		return input.Text()
	}

	return ""
}

func (input *Input) ReadStrings(size int) []string {
	input.Split(bufio.ScanWords)

	strings := make([]string, size)

	for i := 0; i < size && input.Scan(); i++ {
		strings[i] = input.Text()
	}

	return strings
}
