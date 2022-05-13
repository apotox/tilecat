package main

import (
	"os"
	"path/filepath"
)

func listImages(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil

	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
