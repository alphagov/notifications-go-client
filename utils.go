package notify

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func jsonResponse(body io.ReadCloser, obj interface{}) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(b), &obj)

	return err
}
