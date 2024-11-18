package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

var (
	insensitive = flag.Bool("insensitive", false, "Insensitive mode")
)

func main() {}

type Match struct {
	FileName string
	Lines    []string
}

func MatchFiles(fileSystem fs.FS, files []string, candidate string) ([]Match, error) {
	var result []Match

	for _, filename := range files {
		match, err := processFile(fileSystem, filename, candidate)
		if err != nil {
			return nil, err
		}
		result = append(result, match)
	}

	return result, nil
}

func processFile(fileSystem fs.FS, filename string, candidate string) (Match, error) {
	file, err := fileSystem.Open(filename)
	if err != nil {
		return Match{}, fmt.Errorf("error opening file: %s: %w", filename, err)
	}
	defer file.Close()
	fileData, err := io.ReadAll(file)
	if err != nil {
		return Match{}, fmt.Errorf("error reading file: %s: %w", filename, err)
	}

	var match Match
	match.FileName = filename

	match.Lines = splitLines(string(fileData))

	match.Lines = matchLines(&match, candidate)

	return match, nil
}

func splitLines(data string) []string {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(data))

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func matchLines(match *Match, candidate string) []string {
	var result []string

	for _, line := range match.Lines {
		if strings.Contains(line, candidate) {
			result = append(result, line)
		}
	}

	return result
}
