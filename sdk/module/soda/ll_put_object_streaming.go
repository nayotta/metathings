package metathings_module_soda_sdk

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	component "github.com/nayotta/metathings/pkg/component"
)

const (
	C_OBJECT_STREAM_ACTION_PUT_OBJECT_STREAMING = "/v1/actions/put_object_streaming"
)

func (cli *sodaClient) LLPutObjectStreaming(name string, length int64, sha1sum string) (string, error) {
	logger := cli.GetLogger().WithFields(logrus.Fields{
		"#method": "LLPutObjectStreaming",
		"name":    name,
		"length":  length,
		"sha1sum": sha1sum,
	})

	url := cli.joinPath(C_OBJECT_STREAM_ACTION_PUT_OBJECT_STREAMING)
	body := map[string]any{
		"object": map[string]any{
			"name":    name,
			"length":  length,
			"sha1sum": sha1sum,
		},
	}

	buf, err := json.Marshal(body)
	if err != nil {
		logger.WithError(err).Debugf("failed to marshal request body to json string")
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(buf))
	if err != nil {
		logger.WithError(err).Debugf("failed to new request")
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := cli.httpClient.Do(req)
	if err != nil {
		logger.WithError(err).Debugf("failed to do request")
		return "", err
	}
	defer res.Body.Close()

	if ok, err := cli.assertResponse(res, http.StatusNoContent); !ok {
		logger.WithError(err).Debugf("failed to put object streaming")
		return "", err
	}

	osName := res.Header.Get(component.HTTP_SODA_OBJECT_STREAM_NAME)

	logger.WithField("object_stream", osName).Tracef("put object streaming")

	return osName, nil
}
