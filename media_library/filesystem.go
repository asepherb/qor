package media_library

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileSystem struct {
	Base
}

func (f FileSystem) GetFullPath(url string, option Option) (path string, err error) {
	if option["path"] != "" {
		path = filepath.Join(option["path"], url)
	} else {
		path = filepath.Join("./public", url)
	}

	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	}

	return
}

func (f FileSystem) Store(url string, option Option, fileHeader *multipart.FileHeader) error {
	if fullpath, err := f.GetFullPath(url, option); err == nil {
		if dst, err := os.Create(fullpath); err == nil {
			if file, err := fileHeader.Open(); err == nil {
				_, err := io.Copy(dst, file)
				return err
			} else {
				return err
			}
		} else {
			return err
		}
	} else {
		return err
	}
}

func (f FileSystem) Receive(path string) (*os.File, error) {
	return os.Open(path)
}
