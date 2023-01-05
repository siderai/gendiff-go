package gendiff

import (
	"os"
	"path/filepath"
)

func validate_path(path string) (validated_path string, err error) {
	return path, nil
}

func main() {
	args := os.Args
	// TODO: Validate paths
	path_first, err := validate_path(filepath.Base(args[1]))
	if err != nil {
		return nil, err
	}
	path_second, err := validate_path(filepath.Base(args[2]))
	if err != nil {
		return nil, err
	}
}
