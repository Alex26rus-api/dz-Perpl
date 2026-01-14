package files

import (
	"os"
	"path/filepath"
	"strings"
)

func IsJson(name string) bool {
	ext := filepath.Ext(name)
	ext = strings.ToLower(ext)
	if ext == ".json" {
		return true
	}
	return false
}

func Read(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Write(name string, data []byte) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
