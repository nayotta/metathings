package metathings_data_storage_sdk

import (
	"context"
	"encoding/base64"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/stretchr/objx"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type Influxdb2DataStorageOption struct {
	Address  string
	Token    string
	Username string
	Password string
	Org      string
	Bucket   string

	QueryPageSize        int32
	QueryRangeFromOffset time.Duration
}

func NewInfluxdb2DataStorageOption() *Influxdb2DataStorageOption {
	return &Influxdb2DataStorageOption{
		QueryPageSize:        50,
		QueryRangeFromOffset: -24 * time.Hour,
	}
}

type Influxdb2DataStorage struct {
	opt    *Influxdb2DataStorageOption
	influx influxdb2.Client
	logger log.FieldLogger
}

func (s *Influxdb2DataStorage) get_logger() log.FieldLogger {
	return s.logger
}

func (s *Influxdb2DataStorage) Write(ctx context.Context, measurement string, tags map[string]string, data map[string]interface{}) error {
	// TODO(Peer): we should save data into difference buckets

	var ts time.Time
	tsi := ctx.Value("timestamp")
	if tsi != nil {
		ts = tsi.(time.Time)
	} else {
		ts = time.Now()
	}

	logger := s.get_logger().WithFields(log.Fields{
		"measurement": measurement,
		"tags":        tags,
	})

	writer := s.influx.WriteAPIBlocking(s.opt.Org, s.opt.Bucket)
	point := influxdb2.NewPoint(measurement, tags, data, ts)
	if err := writer.WritePoint(ctx, point); err != nil {
		logger.WithError(err).Debugf("failed to write data to influxdb")
		return err
	}

	logger.Debugf("write data")

	return nil
}

func (s *Influxdb2DataStorage) new_query_option() map[string]interface{} {
	return map[string]interface{}{
		"page_size":  s.opt.QueryPageSize,
		"range_from": time.Now().Add(s.opt.QueryRangeFromOffset),
	}
}

func (s *Influxdb2DataStorage) Query(ctx context.Context, measurement string, tags map[string]string, opts ...QueryOption) (QueryResult, error) {
	logger := s.get_logger().WithFields(log.Fields{
		"measurement": measurement,
		"tags":        tags,
	})

	var range_from_str string
	var range_to_str string

	o := s.new_query_option()
	for _, opt := range opts {
		opt(o)
	}
	ox := objx.New(o)

	range_from_str = cast.ToTime(ox.Get("range_from").Data()).Format(time.RFC3339Nano)
	if range_to_if := ox.Get("range_to").Data(); range_to_if != nil {
		range_to_str = cast.ToTime(range_to_if).Format(time.RFC3339Nano)
	}
	query_string_str := cast.ToString(ox.Get("query_string").Data())
	page_size_i32 := cast.ToInt32(ox.Get("page_size").Data())
	logger = logger.WithFields(log.Fields{
		"page_size":  page_size_i32,
		"range_from": range_from_str,
		"range_to":   range_to_str,
	})

	page_token_str := cast.ToString(ox.Get("page_token").Data())
	if page_token_str != "" {
		buf, err := base64.StdEncoding.DecodeString(page_token_str)
		if err != nil {
			logger.WithError(err).Debugf("failed to decode page token")
			return nil, err
		}
		range_from_str = string(buf)
	}

	var sb strings.Builder
	sb.WriteString(`from(bucket: "`)
	sb.WriteString(s.opt.Bucket)
	sb.WriteString(`")`)
	sb.WriteString(` |> range(`)
	sb.WriteString(`start: `)
	sb.WriteString(range_from_str)
	if range_to_str != "" {
		sb.WriteString(`, stop: `)
		sb.WriteString(range_to_str)
	}
	sb.WriteString(`)`)
	sb.WriteString(` |> filter(fn: (r) => r["_measurement"] == "`)
	sb.WriteString(measurement)
	sb.WriteString(`")`)
	for tk, tv := range tags {
		sb.WriteString(` |> filter(fn: (r) => r["`)
		sb.WriteString(tk)
		sb.WriteString(`"] == "`)
		sb.WriteString(tv)
		sb.WriteString(`")`)
	}
	sb.WriteString(query_string_str)
	if page_size_i32 > 0 {
		sb.WriteString(` |> limit(n: `)
		sb.WriteString(strconv.Itoa(int(page_size_i32)))
		sb.WriteString(`)`)
	}

	query_str := sb.String()
	logger.WithField("full_query_string", query_str).Debugf("query string")

	querier := s.influx.QueryAPI(s.opt.Org)
	qtr, err := querier.Query(ctx, query_str)
	if err != nil {
		logger.WithError(err).Debugf("failed to query in influxdb")
		return nil, err
	}

	result_string_map_index_by_time := map[time.Time]map[string]interface{}{}
	for qtr.Next() {
		rc := qtr.Record()
		rc_ts := rc.Time()
		vals, ok := result_string_map_index_by_time[rc_ts]
		if !ok {
			vals = map[string]interface{}{}
		}
		vals[rc.Field()] = rc.Value()
		result_string_map_index_by_time[rc_ts] = vals
	}

	if len(result_string_map_index_by_time) == 0 {
		logger.Debugf("end of query token")
		return NewQueryResult(nil, ""), io.EOF
	}

	var qrs []QueryRecord
	for tm, dat := range result_string_map_index_by_time {
		qrs = append(qrs, NewQueryRecord(tm, dat))
	}

	sort.Slice(qrs, func(i, j int) bool { return qrs[i].Time().Sub(qrs[j].Time()) < 0 })

	next_page_token_str := qrs[len(qrs)-1].Time().Add(time.Nanosecond).Format(time.RFC3339Nano)
	next_page_token_str = base64.StdEncoding.EncodeToString([]byte(next_page_token_str))

	return NewQueryResult(qrs, next_page_token_str), nil
}

func NewInfluxdb2DataStorage(args ...interface{}) (DataStorage, error) {
	var logger log.FieldLogger
	opt := NewInfluxdb2DataStorageOption()

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"address":  opt_helper.ToString(&opt.Address),
		"token":    opt_helper.ToString(&opt.Token),
		"username": opt_helper.ToString(&opt.Username),
		"password": opt_helper.ToString(&opt.Password),
		"org":      opt_helper.ToString(&opt.Org),
		"bucket":   opt_helper.ToString(&opt.Bucket),
		"logger":   opt_helper.ToLogger(&logger),
	})(args...); err != nil {
		return nil, err
	}

	// TODO(Peer): allow userpass login
	influx := influxdb2.NewClient(opt.Address, opt.Token)

	s := &Influxdb2DataStorage{
		opt:    opt,
		influx: influx,
		logger: logger,
	}

	logger.Debugf("new influxdb client")

	return s, nil
}

func init() {
	registry_data_storage_factory("influxdb2", NewInfluxdb2DataStorage)
}
