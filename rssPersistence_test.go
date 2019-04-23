package rsspersistence_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/vichar/rsspersistence"
)

func TestReadingConfiguration(t *testing.T) {

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

	t.Run("ParseContent File to Object", func(t *testing.T) {
		data := `
        key: 0c681e1e-48ad-11e9-8646-d663bd873d93
        topic: /topics/test
        `
		content, error := rsspersistence.ParseContent(data)
		if error != nil {
			get := error.Error()
			t.Errorf("It should not return any errors %s", get)
		} else {

			if "0c681e1e-48ad-11e9-8646-d663bd873d93" != content.Key {
				t.Errorf("It should contain assigned key: 0c681e1e-48ad-11e9-8646-d663bd873d9")
			}
			if "/topics/test" != content.Topic {
				t.Errorf("It should contain assigned topic: /topics/test")
			}
		}
	})

	t.Run("Connecting to Firebase Database", func(t *testing.T) {
		result, error := rsspersistence.ConnectFirebase("fbKey", "topic")
		if error != nil {
			get := error.Error()
			t.Errorf("It should not return any errors %s", get)
		} else {
			if !result {
				t.Errorf("It should not return any errors %s", strconv.FormatBool(result))
			}
		}
	})
}
