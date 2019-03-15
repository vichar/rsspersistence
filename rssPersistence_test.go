package rsspersistence_test

import (
	"strings"
	"testing"

	"github.com/vichar/rsspersistence"
)

func TestParseDatabaseConfiguration(t *testing.T) {

	t.Run("ReadingTextFile Happy Path", func(t *testing.T) {
		content, error := rsspersistence.ReadTextFile("./test.txt")
		if error != nil {
			t.Errorf("Test File should be present")
		} else {
			get := strings.TrimSpace(content)
			want := "Hello World!"
			if want != get {
				t.Errorf("It should say %s but get %s", want, get)
			}

		}
	})

	t.Run("ReadingTextFile File Not Exist", func(t *testing.T) {
		content, error := rsspersistence.ReadTextFile("")
		if error != nil {
			get := error.Error()
			want := "open : no such file or directory"
			if want != get {
				t.Errorf("It should say %s but get %s", want, get)
			}
		} else {
			get := strings.TrimSpace(content)
			want := "open : no such file or directory"
			t.Errorf("Expected to get %s but get %s", want, get)
		}
	})
}
