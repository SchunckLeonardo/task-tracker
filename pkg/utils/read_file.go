package utils

import (
	"encoding/json"
	"os"
)

func ReadFile[T any](filename string) (*T, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var value T
	err = json.Unmarshal(file, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}
