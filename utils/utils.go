package utils

import (
	"os"
	"strings"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func IngestInputFile(path string) []string {

	content, err := os.ReadFile(path)

	CheckError(err)

	strContent := string(content)
	lines := strings.Split(strContent, "\n")

	return lines
}
