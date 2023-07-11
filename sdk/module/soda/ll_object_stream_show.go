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
	C_OBJECT_STREAM_ACTION_SHOW = "/v1/object_streams/%s/actions/show"
)

type ObjectStreamInfo struct {
	Sha1sum     string
	Length      int64
	Uploaded    int64
	MaxAge      time.Duration
	Remained    time.Duration
	Offset      int64
	ChunkLength int64
}

func (cli *sodaClient) LLObjectStreamShow(name string) (info ObjectStreamInfo, err error) {
	logger := cli.GetLogger().WithFields(logrus.Fields{
		"#method": "LLObjectStreamShow",
		"name":    name,
	})

	url := cli.joinPath(fmt.Sprintf(C_OBJECT_STREAM_ACTION_SHOW, name))
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader("{}"))
	if err != nil {
		logger.WithError(err).Debugf("failed to new request")
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := cli.httpClient.Do(req)
	if err != nil {
		logger.WithError(err).Debugf("failed to do http request")
		return
	}
	defer res.Body.Close()

	ok, err := cli.assertResponse(res, http.StatusNoContent)
	if !ok {
		logger.WithError(err).Debugf("failed to show")
		return
	}

	info.Sha1sum = res.Header.Get(component.HTTP_SODA_OBJECT_SHA1SUM)
	info.Length = cast.ToInt64(res.Header.Get(component.HTTP_SODA_OBJECT_LENGTH))
	info.Uploaded = cast.ToInt64(res.Header.Get(component.HTTP_SODA_OBJECT_UPLOADED_LENGTH))
	info.MaxAge = cast.ToDuration(res.Header.Get(component.HTTP_SODA_OBJECT_STREAM_MAX_AGE))
	info.Remained = cast.ToDuration(res.Header.Get(component.HTTP_SODA_OBJECT_STREAM_REMAINED))
	info.Offset = cast.ToInt64(res.Header.Get(component.HTTP_SODA_OBJECT_STREAM_CHUNK_OFFSET))
	info.ChunkLength = cast.ToInt64(res.Header.Get(component.HTTP_SODA_OBJECT_STREAM_CHUNK_LENGTH))

	logger.Tracef("show")

	return
}
