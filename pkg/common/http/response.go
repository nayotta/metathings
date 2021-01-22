package http_helper

import (
	"encoding/json"
	"net/http"
)

type JSONResponseWriter interface {
	http.ResponseWriter
	WriteJSON(interface{}) error
	WriteJSONString(string) error
}

type jsonResponseWriter struct {
	http.ResponseWriter
}

func (w *jsonResponseWriter) WriteJSON(data interface{}) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = w.WriteJSONString(string(buf))
	if err != nil {
		return err
	}

	return nil
}

func (w *jsonResponseWriter) WriteJSONString(s string) error {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(s))
	if err != nil {
		return err
	}

	return nil
}

func WrapJSONResponseWriter(w http.ResponseWriter) JSONResponseWriter {
	return &jsonResponseWriter{w}
}

func ConvertError(err error) interface{} {
	return map[string]interface{}{
		"error": err.Error(),
	}
}
