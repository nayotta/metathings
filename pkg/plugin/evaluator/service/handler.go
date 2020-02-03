package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var err error
	var buf []byte

	dat := make(map[string]interface{})
	dev_id := r.Header.Get("MT-Evaluator-Device")
	flw_nm := r.Header.Get("MT-Evaluator-Flow")
	flwst_id := r.Header.Get("MT-Operator-FlowSet")

	w.Header().Set("Content-Type", "application/json")

	if buf, err = ioutil.ReadAll(r.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request"}`))
		return
	}

	if err = json.Unmarshal(buf, &dat); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "bad request"}`))
		return
	}

	res := make(map[string]interface{})
	res["DeviceID"] = dev_id
	res["FlowName"] = flw_nm
	res["FlowSetID"] = flwst_id
	res["Data"] = dat

	if buf, err = json.Marshal(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "internal server error"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
