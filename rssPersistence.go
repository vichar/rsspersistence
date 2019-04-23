package rsspersistence

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ReadTextFile takes path and returns content of file
func ReadTextFile(path string) (string, error) {
	buffer, error := ioutil.ReadFile(path)
	return string(buffer), error
}

// ParseContent parsing content of a YML files
func ParseContent(content string) (FirebaseDatabaseConnnection, error) {
	connection := FirebaseDatabaseConnnection{}
	error := yaml.Unmarshal([]byte(content), &connection)
	return connection, error
}

// ConnectFirebase connecting to Firebase Database
func ConnectFirebase(firebaseKey string, topic string) (bool, error) {

	return true, nil
}

// FirebaseDatabaseConnnection contains data for connecting to Fireware Database
type FirebaseDatabaseConnnection struct {
	Key   string
	Topic string
}
