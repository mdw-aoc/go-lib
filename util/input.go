package util

import (
	"bytes"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/mdw-aoc/go-lib/parse"
	"github.com/mdw-aoc/inputs/v2/inputs"
)

func InputBytes() []byte {
	_, path, _, _ := runtime.Caller(0)
	this := path
	for caller := 1; path == this; caller++ {
		_, path, _, _ = runtime.Caller(caller)
	}
	pattern := regexp.MustCompile(`advent-of-code/go/(\d{4})/day(\d{2})/`)
	matches := pattern.FindStringSubmatch(path)
	year, day := matches[1], matches[2]
	for strings.Contains(path, "advent-of-code") {
		path = filepath.Dir(path)
	}
	return inputs.Read(parse.Int(year), parse.Int(day)).Bytes()
}

func InputInt() int {
	i, _ := strconv.Atoi(InputString())
	return i
}

func InputInts(sep string) []int {
	return parse.Ints(strings.Split(InputString(), sep))
}

func InputString() string {
	return strings.TrimSpace(string(InputBytes()))
}

func InputLines() []string {
	return strings.Split(InputString(), "\n")
}

func InputScanner() *Scanner {
	return NewScanner(bytes.NewReader(InputBytes()))
}
