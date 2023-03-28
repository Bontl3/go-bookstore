package utils

// the data will be in JSON and we need to unmarshel it so that the app can use it

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// function will receive a json body and then it will need to unbudle it
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}

}
