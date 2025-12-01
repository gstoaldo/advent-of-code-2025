package utils

import (
	"os"
	"strconv"
	"strings"
)

func FilePath() string {
	filepath := "input.txt"

	if len(os.Args) == 2 {
		filepath = os.Args[1]
	}

	return filepath
}

func ReadFile(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic("failed to open file")
	}

	return strings.TrimSpace(string(file))
}

func ReadLines(path string) []string {
	file := ReadFile(path)

	return strings.Split(file, "\n")
}

func ToInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic("failed to convert string to int")
	}

	return v
}
