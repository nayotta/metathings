package metathings_module_soda_sdk

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"

	component "github.com/nayotta/metathings/pkg/component"
)

const (
	C_OBJECT_STREAM_ACTION_WRITE_CHUNK = "/v1/object_streams/%s/actions/write_chunk"
)

func (cli *sodaClient) LLObjectStreamWriteChunk(name string, offset int64, length int64, chunk []byte) error {
	logger := cli.GetLogger().WithFields(logrus.Fields{
		"#method": "LLObjectStreamWriteChunk",
		"name":    name,
		"offset":  offset,
		"length":  length,
	})

	url := cli.joinPath(fmt.Sprintf(C_OBJECT_STREAM_ACTION_WRITE_CHUNK, name))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(chunk))
	if err != nil {
		logger.WithError(err).Debugf("failed to new request")
		return err
	}

	h := sha1.New()
	h.Write(chunk)
	sha1sum := hex.EncodeToString(h.Sum(chunk))

	req.Header.Add("Content-Type", "application/octet-stream")
	req.Header.Add(component.HTTP_SODA_OBJECT_STREAM_CHUNK_OFFSET, cast.ToString(offset))
	req.Header.Add(component.HTTP_SODA_OBJECT_STREAM_CHUNK_LENGTH, cast.ToString(length))
	req.Header.Add(component.HTTP_SODA_OBJECT_STREAM_CHUNK_SHA1SUM, sha1sum)

	res, err := cli.httpClient.Do(req)
	if err != nil {
		logger.WithError(err).Debugf("failed to do http request")
		return err
	}
	defer res.Body.Close()

	if ok, err := cli.assertResponse(res, http.StatusNoContent); !ok {
		logger.WithError(err).Debugf("failed to write chunk")
		return err
	}

	logger.Tracef("write chunk")

	return nil
}
