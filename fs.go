package forge

import (
	"io"
	"os"
	"path/filepath"
)

var fs FileSystem = new(NativeFS)

// FileSystem interface to expand configuration file parsing possiblities.
type FileSystem interface {
	Open(filename string) (io.Reader, error)
	Glob(pattern string) ([]string, error)
}

// RegisterFS method to register new file system into forge.
func RegisterFS(fileSystem FileSystem) {
	fs = fileSystem
}

// NativeFS type defines methods for native file system.
type NativeFS uint8

func (NativeFS) Open(filename string) (io.Reader, error) {
	return os.Open(filename)
}

func (NativeFS) Glob(pattern string) ([]string, error) {
	return filepath.Glob(pattern)
}

// Close method closes the give file if its compliant to `io.Closer`
func close(f interface{}) error {
	if c, ok := f.(io.Closer); ok {
		return c.Close()
	}
	return nil
}
