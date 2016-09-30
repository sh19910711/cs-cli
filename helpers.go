package main
import (
	"os"
	"io/ioutil"
	"github.com/parnurzeal/gorequest"
)

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func DownloadFile(url, path string) []error {
	_, body, errs := gorequest.New().Get(url).EndBytes()
	ioutil.WriteFile(path, body, 0644)
	return errs
}
