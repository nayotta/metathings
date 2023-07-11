package metathings_module_soda_sdk

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"

	component "github.com/nayotta/metathings/pkg/component"
)

const (
	C_OBJECT_STREAM_ACTION_NEXT_CHUNK = "/v1/object_streams/%s/actions/next_chunk"
)

func (cli *sodaClient) LLObjectStreamNextChunk(name string) (remained time.Duration, offset int64, length int64, err error) {
	logger := cli.GetLogger().WithFields(logrus.Fields{
		"#method": "LLObjectStreamNextChunk",
		"name":    name,
	})

	url := cli.joinPath(fmt.Sprintf(C_OBJECT_STREAM_ACTION_NEXT_CHUNK, name))
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader("{}"))
	if err != nil {
		logger.WithError(err).Debugf("failed to new request")
		return 0, 0, 0, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := cli.httpClient.Do(req)
	if err != nil {
		logger.WithError(err).Debugf("failed to do http request")
		return 0, 0, 0, err
	}
	defer res.Body.Close()

	if ok, err := cli.assertResponse(res, http.StatusNoContent); !ok {
		logger.WithError(err).Debugf("failed to next chunk")
		return 0, 0, 0, err
	}

	remained = cast.ToDuration(res.Header.Get(component.HTTP_SODA_OBJECT_STREAM_REMAINED))
	offset = cast.ToInt64(res.Header.Get(component.HTTP_SODA_OBJECT_STREAM_CHUNK_OFFSET))
	length = cast.ToInt64(res.Header.Get(component.HTTP_SODA_OBJECT_STREAM_CHUNK_LENGTH))

	logger.Tracef("next chunk")

	return
}
