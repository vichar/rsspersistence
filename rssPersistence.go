package rsspersistence

import (
	"io/ioutil"
)

// ReadTextFile takes path and returns content of file
func ReadTextFile(path string) (string, error) {
	buffer, error := ioutil.ReadFile(path)
	return string(buffer), error
}
