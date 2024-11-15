package main

import "io/fs"

func main() {}

type Match struct {
}

func MatchFiles(fileSystem fs.FS, files []string, candidate string) []Match {
	return []Match{}
}
