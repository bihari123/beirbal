package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

// if isFile is true then it will give only files, else only folders
func GetDirChildren(filePath string) (filesNFolders []fs.FileInfo) {
	filesNFolders, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for index, f := range filesNFolders {
		if f.IsDir() {
			// eliminate the folders, only keep the file
			filesNFolders = append(filesNFolders[:index], filesNFolders[index+1:]...)
		}
	}
	return
}

func MakeDirIfNotExists(filePath string) (err error) {
	if !Exists(filePath) {
		if err = os.MkdirAll(filePath, os.ModePerm); err != nil {
			err = errors.New(fmt.Sprintf("Problem making directory %v: %v", filePath, err))
			return
		}
	}
	return
}

func WriteAppendFile(filePath, fileContent string) (err error) {
	var idx int64

	if Exists(filePath) {
		content, err := GetFileContent(filePath)
		if err != nil {
			log.Fatal(err)
		}

		data := []byte(content)
		idx = int64(len(data))

	}

	var f *os.File

	f, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		return
	}

	defer f.Close()

	newEntry := fmt.Sprintf("%v\n", fileContent)

	data2 := []byte(newEntry)
	_, err = f.WriteAt(data2, idx)

	return
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	} else {
		log.Fatal(fmt.Sprintf("some error happened at checking the existence of the directory: %s", path), err)
	}

	return false
}

func GetFileContent(filePath string) (content string, err error) {
	contentBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	content = string(contentBytes)

	return
}
