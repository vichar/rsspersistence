package rsspersistence

import (
	"testing"

	"github.com/vichar/rsspersistence"
)

func TestParseDatabaseConfiguration(t *testing.T) {

	t.Run("ReadingTextFile", func(t *testing.T) {
		rsspersistence.ReadTextFile("test")
	})
}

// func TestParseHTTPResponseWithValidURL(t *testing.T) {
// 	t.Run("ParseHTTPResponse", func(t *testing.T) {
// 		response, error := rssfeed.HTTPGet("https://www.npr.org/rss/podcast.php?id=500005")
// 		if error == nil {
// 			var getResponse rssfeed.HTTPResponse = rssfeed.ParseHTTPResponse(response)
// 			if getResponse.HTTPStatus != "200 OK" {
// 				t.Errorf("HTTP Status shoud be %s but gets %s", getResponse.HTTPStatus, "200 OK")
// 			}
// 			if getResponse.HTTPStatusCode != 200 {
// 				t.Errorf("HTTP Status Code should be %s but gets %s", strconv.Itoa(getResponse.HTTPStatusCode), "200")
// 			}
// 			if len(getResponse.Body) <= 0 {
// 				t.Errorf("HTTP Response Body should not be empty but length is %s", strconv.Itoa(len(getResponse.Body)))
// 			}
// 		} else {
// 			t.Errorf("HTTPGet should not return error for a valid URL but gets %s", error.Error())
// 		}
// 	})
// }
//
// func TestHTTPGetWithValidURL(t *testing.T) {
// 	t.Run("HTTPGet should return valid String URL", func(t *testing.T) {
// 		response, error := rssfeed.HTTPGet("https://www.npr.org/rss/podcast.php?id=500005")
// 		want := "200 OK"
// 		if want != response.Status {
// 			t.Errorf("It should say %s but gets %s", want, response.Status)
// 		}
// 		if error != nil {
// 			t.Errorf("An Unexpected Error occurs %s", error.Error())
// 		}
//
// 	})
//
// }
//
// func TestGetContentWithInvalidURL(t *testing.T) {
// 	t.Run("HTTPGet should return Empty String", func(t *testing.T) {
// 		response, error := rssfeed.HTTPGet("")
// 		if error == nil {
// 			t.Errorf("Invalid URL is Expected %s %s", response.Status, error.Error())
// 		}
// 		if error.Error() != "Invalid URL" {
// 			t.Errorf("Unexpected Error has occurred %s %s", response.Status, error.Error())
// 		}
// 	})
//
// }
