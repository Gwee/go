package main

import (
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func main() {

	//Usage example
	//winPath := "C:\\temp"
	//linPath := "/Users/guy/Desktop"
	//var buf []byte
	//itemsFound, err := searchWord("test",linPath, &buf)
	//if err != nil {
	//	log.Fatal(err)
	//}
}

func searchWord(term string, path string, buffer *[]byte) (string, error) {
	if len(path) == 0 {
		return "", errors.New("Please enter a valid path")
	}
	if len(term) == 0 {
		return "", errors.New("Please enter a valid search term")
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	//traverse directories by going over each file and if we find a dir enter recursion call
	for _, file := range files {
		path := filepath.Join(path, file.Name())
		result, err := filepath.Abs(path)
		if err != nil {
			log.Fatal(err)
		}
		//if we find a matching file/dir substring, add it to the buffer
		if strings.Contains(file.Name(), term) {
			*buffer = append(*buffer, result+"\n"...)
		}
		//if the file is a dir, enter it recursively
		if file.Mode().IsDir() {
			searchWord(term, path, buffer)
		}
	}
	return string(*buffer), err
}