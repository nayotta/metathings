package metathings_module_soda_sdk

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"mime/multipart"
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

	h := sha1.New()
	h.Write(chunk)
	sha1sumStr := hex.EncodeToString(h.Sum(nil))

	var body bytes.Buffer
	wr := multipart.NewWriter(&body)

	wr.WriteField(component.HTTP_SODA_OBJECT_STREAM_CHUNK_OFFSET, cast.ToString(offset))
	wr.WriteField(component.HTTP_SODA_OBJECT_STREAM_CHUNK_LENGTH, cast.ToString(length))
	wr.WriteField(component.HTTP_SODA_OBJECT_STREAM_CHUNK_SHA1SUM, sha1sumStr)

	part, err := wr.CreateFormFile("file", "chunk")
	if err != nil {
		logger.WithError(err).Debugf("failed to create form file")
		return err
	}
	if _, err = part.Write(chunk); err != nil {
		logger.WithError(err).Debugf("failed to write chunk to part")
		return err
	}

	if err = wr.Close(); err != nil {
		logger.WithError(err).Debugf("failed to close multipart writer")
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, &body)
	if err != nil {
		logger.WithError(err).Debugf("failed to new http request")
		return err
	}

	req.Header.Set("Content-Type", wr.FormDataContentType())

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
