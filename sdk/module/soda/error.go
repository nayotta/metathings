package metathings_module_soda_sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrPutObjectTimeout = fmt.Errorf("put object timeout")
)

func (cli *sodaClient) parseResponseError(res *http.Response) (error, error) {
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resBody := make(map[string]string)
	if err = json.Unmarshal(buf, &resBody); err != nil {
		return nil, fmt.Errorf("%w: %s", err, string(buf))
	}

	errStr := resBody["error"]
	return fmt.Errorf(errStr), nil
}

func (cli *sodaClient) assertResponse(res *http.Response, expectCode int) (bool, error) {
	if res.StatusCode != expectCode {
		resErr, err := cli.parseResponseError(res)
		if err != nil {
			return false, err
		}

		return false, resErr
	}

	return true, nil
}
