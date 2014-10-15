package file_system

import (
	"io/ioutil"
	"os"
)

type FileStore struct {
	FilePath string
}

func (this *FileStore) CreateFile(path string, content []byte) (err error) {
	err = ioutil.WriteFile(path, content, os.ModePerm)
	return
}

func (this *FileStore) UpdateFile(path string, content []byte) (err error) {
	return
}

func (this *FileStore) CreateDir(path string) (err error) {
	return
}

func (this *FileStore) ReadFile(path string) (content []byte, err error) {
	content, err = ioutil.ReadFile(path)
	return
}

func (this *FileStore) ReadDir(path string) (paths []string, err error) {
	return
}

func (this *FileStore) Move(path string, newPath string) (err error) {
	return
}

func (this *FileStore) Delete(path string) (err error) {
	return
}
