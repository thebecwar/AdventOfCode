package loader

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

type Loader struct {
	Filename string
	Lines    []string
}

func NewLoader(filename string) (*Loader, error) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("failed to get source location")
	}

	path := filepath.Dir(file)
	filename = filepath.Join(path, "..", "data", filename)

	loader := &Loader{Filename: filename}
	err := loader.Load()
	return loader, err
}

func (l *Loader) Load() error {
	file, err := os.Open(l.Filename)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l.Lines = append(l.Lines, scanner.Text())
	}

	return scanner.Err()
}
