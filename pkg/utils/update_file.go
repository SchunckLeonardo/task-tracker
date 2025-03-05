package utils

import (
	"encoding/json"
	"os"
)

func UpdateFile[T any](filename string, data T) error {
	updatedFile, err := json.Marshal(data)
	if err != nil {
		return err
	}

	jsonFmt := []byte(`{"tasks": ` + string(updatedFile) + `}`)

	err = os.WriteFile(filename, jsonFmt, 0666)

	return err
}
