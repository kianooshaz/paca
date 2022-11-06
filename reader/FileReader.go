package reader

import (
	"os"
)

func FileReader(path string) (Source, error) {
	input, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	sourceString := string(input)
	return Source(sourceString), nil
}
