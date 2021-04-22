package wgm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type ApiRequestError struct {
	message string `json:message`
}

func (a *ApiRequestError) Error(s string) {
	a.message = s
}

func (a ApiRequestError) Marshal() []byte {
	buf, _ := json.Marshal(a)
	return buf
}

func mainApiHandler(w http.ResponseWriter, r *http.Request) {
	var auth string
	defer r.Body.Close()
	reqErr := new(ApiRequestError)
	// Check if mimetype is json and auth header exists
	if r.Header.Get("Content-Type") == "application/json" || r.Header.Get("content-type") == "application/json" {
		// Only assign new value if it was null (so we can check both Authorization and authorization)
		assignIfNull(&auth, r.Header.Get("Authorization"))
		assignIfNull(&auth, r.Header.Get("authorization"))
		if auth == "" {
			reqErr.Error("Incorrect authorization header")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(reqErr.Marshal())
			return
		}
		// Compile token regexp
		matcher, err := regexp.Compile("/^Bearer (.*)$/")
		if err != nil {
			panic("Could not compile authorization regular expression")
		}
		// Find token in regexp
		match := matcher.FindString(auth)
		if match == "" {
			reqErr.Error("Could not find any proper bearer token, though authorization header was recognized")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(reqErr.Marshal())
			return
		}
		// Assign back to auth
		auth = match
	} else {
		reqErr.Error("Improper request mime type, expecting application/json")
		// Return an unsupported media type error (415)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write(reqErr.Marshal())
		return
	}

	// Create new buffer to store response body
	buf := new(bytes.Buffer)
	// Read io.ReadCloser and store it in previous byte buffer allocation
	buf.ReadFrom(r.Body)
}

func main() {
	fmt.Println("This is a test")
	http.Handle("/", mainApiHandler)
}
