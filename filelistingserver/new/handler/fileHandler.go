package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

func FileServerHandler(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		//http.Error(writer, "Invalid Path", http.StatusBadRequest)
		return errors.New("path must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):]

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	writer.Write(all)

	return nil
}
