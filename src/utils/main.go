//The main.go file above contains utility functions such as the AbsPath function, which gets the absolute path to a specified directory; the ReadFiles function, which reads a directory recursively and gets the absolute path to its files and subfolder files, and the RootDir function, which gets the path to the root directory of the project.

package utils

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func AbsPath(path string) string {
	r := RootDir()
	return filepath.Join(r, path)
}

func ReadFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
