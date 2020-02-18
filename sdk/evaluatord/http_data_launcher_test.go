package metathings_evaluatord_sdk

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	test_source_id   = "test_source_id"
	test_source_type = "test_source_type"
	test_data        = []byte(`{}`)
)

type HttpDataLauncherTestSuite struct {
	suite.Suite
}

func (ts *HttpDataLauncherTestSuite) TestLaunch() {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ts.Equal("POST", r.Method)
		ts.Equal("application/json", r.Header.Get("Content-Type"))
		ts.Equal(test_source_id, r.Header.Get(HTTP_HEADER_SOURCE_ID))
		ts.Equal(test_source_type, r.Header.Get(HTTP_HEADER_SOURCE_TYPE))
		ts.Equal("json", r.Header.Get(HTTP_HEADER_DATA_ENCODER))
		buf, err := ioutil.ReadAll(r.Body)
		ts.Nil(err)
		ts.Equal(test_data, buf)
	}))
	defer s.Close()

	dl, err := NewDataLauncher("http", "endpoint", s.URL, "data_encoder", "json")
	ts.Nil(err)

	dat, err := DataFromBytes(test_data)
	ts.Nil(err)

	err = dl.Launch(context.TODO(), NewResource(test_source_id, test_source_type), dat)
	ts.Nil(err)
}

func TestHttpDataLauncherTestSuite(t *testing.T) {
	suite.Run(t, new(HttpDataLauncherTestSuite))
}
