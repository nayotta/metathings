package metathings_data_storage_sdk

import (
	"context"
	"sync"
	"time"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type QueryRecord interface {
	Time() time.Time
	Data() map[string]interface{}
}

type queryRecord struct {
	time time.Time
	data map[string]interface{}
}

func (rec *queryRecord) Time() time.Time {
	return rec.time
}

func (rec *queryRecord) Data() map[string]interface{} {
	return rec.data
}

func NewQueryRecord(tm time.Time, dat map[string]interface{}) QueryRecord {
	return &queryRecord{
		time: tm,
		data: dat,
	}
}

type QueryResult interface {
	Records() []QueryRecord
	NextPageToken() string
}

type queryResult struct {
	records         []QueryRecord
	next_page_token string
}

func (ret *queryResult) Records() []QueryRecord {
	return ret.records
}

func (ret *queryResult) NextPageToken() string {
	return ret.next_page_token
}

func NewQueryResult(recs []QueryRecord, npt string) QueryResult {
	return &queryResult{
		records:         recs,
		next_page_token: npt,
	}
}

type QueryOption func(map[string]interface{})

func RangeFrom(t time.Time) QueryOption {
	return func(o map[string]interface{}) {
		o["range_from"] = t
	}
}

func RangeTo(t time.Time) QueryOption {
	return func(o map[string]interface{}) {
		o["range_to"] = t
	}
}

func QueryString(s string) QueryOption {
	return func(o map[string]interface{}) {
		o["query_string"] = s
	}
}

func PageSize(size int32) QueryOption {
	return func(o map[string]interface{}) {
		o["page_size"] = size
	}
}

func PageToken(token string) QueryOption {
	return func(o map[string]interface{}) {
		o["page_token"] = token
	}
}

type DataStorage interface {
	Write(ctx context.Context, measurement string, tags map[string]string, data map[string]interface{}) error
	Query(ctx context.Context, measurement string, tags map[string]string, opts ...QueryOption) (QueryResult, error)
}

type DataStorageFactory func(...interface{}) (DataStorage, error)

var data_storage_factories map[string]DataStorageFactory
var data_storage_factories_once sync.Once

func registry_data_storage_factory(name string, fty DataStorageFactory) {
	data_storage_factories_once.Do(func() {
		data_storage_factories = map[string]DataStorageFactory{}
	})
	data_storage_factories[name] = fty
}

func NewDataStorage(name string, args ...interface{}) (DataStorage, error) {
	fty, ok := data_storage_factories[name]
	if !ok {
		return nil, ErrUnsupportedDataStorageDriver
	}

	return fty(args...)
}

func ToDataStorage(ds *DataStorage) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool
		if *ds, ok = val.(DataStorage); !ok {
			return opt_helper.InvalidArgument(key)
		}
		return nil
	}
}
