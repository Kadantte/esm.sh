package storage

import (
	"io"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

type localFSDriver struct{}

func (driver *localFSDriver) Open(root string, options url.Values) (FileSystem, error) {
	root = filepath.Clean(root)
	err := ensureDir(root)
	if err != nil {
		return nil, err
	}
	return &localFSLayer{root}, nil
}

type localFSLayer struct {
	root string
}

func (fs *localFSLayer) Stat(name string) (FileStat, error) {
	fullPath := path.Join(fs.root, name)
	fi, err := os.Lstat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return fi, nil
}

func (fs *localFSLayer) Open(name string) (file io.ReadSeekCloser, err error) {
	fullPath := path.Join(fs.root, name)
	file, err = os.Open(fullPath)
	if err != nil && os.IsNotExist(err) {
		err = ErrNotFound
	}
	return
}

func (fs *localFSLayer) WriteFile(name string, content io.Reader) (written int64, err error) {
	fullPath := path.Join(fs.root, name)
	err = ensureDir(path.Dir(fullPath))
	if err != nil {
		return
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return
	}
	defer file.Close()

	written, err = io.Copy(file, content)
	return
}

func (fs *localFSLayer) Remove(name string) (err error) {
	err = os.Remove(path.Join(fs.root, name))
	return
}

func (fs *localFSLayer) RemoveAll(dirname string) (err error) {
	err = os.RemoveAll(path.Join(fs.root, dirname))
	return
}

func ensureDir(dir string) (err error) {
	_, err = os.Lstat(dir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
	}
	return
}

func init() {
	RegisterFileSystem("local", &localFSDriver{})
}
