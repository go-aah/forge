package forge

import (
	"io"
	"os"
	"path/filepath"
)

var (
	nfs FileSystem = new(NativeFS) // Native file system
	efs FileSystem                 // External file system
)

// FileSystem interface to expand configuration file parsing possiblities.
type FileSystem interface {
	Open(filename string) (io.Reader, error)
	Glob(pattern string) ([]string, error)
}

// RegisterFS method to register new file system into forge.
func RegisterFS(fileSystem FileSystem) {
	efs = fileSystem
}

// NativeFS type defines methods for native file system.
type NativeFS uint8

func (NativeFS) Open(filename string) (io.Reader, error) {
	return os.Open(filename)
}

func (NativeFS) Glob(pattern string) ([]string, error) {
	return filepath.Glob(pattern)
}

func open(filename string) (io.Reader, error) {
	if efs == nil || exists(filename) {
		return nfs.Open(filename)
	}
	return efs.Open(filename)
}

func glob(pattern string) ([]string, error) {
	if efs == nil {
		return nfs.Glob(pattern)
	}
	return efs.Glob(pattern)
}

func exists(name string) bool {
	_, err := os.Lstat(name)
	return !os.IsNotExist(err)
}

// Close method closes the give file if its compliant to `io.Closer`
func close(f interface{}) error {
	if c, ok := f.(io.Closer); ok {
		return c.Close()
	}
	return nil
}
