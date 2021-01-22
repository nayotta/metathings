package http_helper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/stretchr/objx"
)

type JSONRequest struct {
	*http.Request
}

func (r *JSONRequest) JSON() (objx.Map, error) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var body map[string]interface{}
	if err = json.Unmarshal(buf, &body); err != nil {
		return nil, err
	}

	bodyx := objx.New(body)

	return bodyx, nil
}

func WrapJSONRequest(r *http.Request) *JSONRequest {
	return &JSONRequest{r}
}
