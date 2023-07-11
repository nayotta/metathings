package metathings_module_soda_sdk

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	C_OBJECT_STREAM_ACTION_CANCEL = "/v1/object_streams/%s/actions/cancel"
)

func (cli *sodaClient) LLObjectStreamCancel(name string) error {
	logger := cli.GetLogger().WithFields(logrus.Fields{
		"#method": "LLObjectStreamCancel",
		"name":    name,
	})

	url := cli.joinPath(fmt.Sprintf(C_OBJECT_STREAM_ACTION_CANCEL, name))
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader("{}"))
	if err != nil {
		logger.WithError(err).Debugf("failed to new request")
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := cli.httpClient.Do(req)
	if err != nil {
		logger.WithError(err).Debugf("failed to do http request")
		return err
	}
	defer res.Body.Close()

	if ok, err := cli.assertResponse(res, http.StatusNoContent); !ok {
		logger.WithError(err).Debugf("failed to cancel")
		return err
	}

	logger.Tracef("cancel object stream")

	return nil
}
