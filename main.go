package main

import (
	"errors"
	"fmt"
)

type Reader interface {
    Read() string
}

type Writer interface {
    Write(data string) error
}

type Storage interface {
    Reader
    Writer
}

type FileStorage struct {
    Name string
}

func (f FileStorage) Read() string {
    return "Reading " + f.Name
}

func (f FileStorage) Write(data string) error {
    return errors.New(data)
}

func processFile(storage Storage) {
    fmt.Println(storage.Read())
    fmt.Println(storage.Write("error"))
}

func main() {
 file_storage := FileStorage{
        Name: "document.txt",
    }


    processFile(file_storage)
}