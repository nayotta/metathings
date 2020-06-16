package metathings_data_storage_sdk

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	log "github.com/sirupsen/logrus"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type Influxdb2DataStorageOption struct {
	Address  string
	Token    string
	Username string
	Password string
	Org      string
	Bucket   string
}

type Influxdb2DataStorage struct {
	opt    *Influxdb2DataStorageOption
	influx influxdb2.InfluxDBClient
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

	writer := s.influx.WriteApiBlocking(s.opt.Org, s.opt.Bucket)
	point := influxdb2.NewPoint(measurement, tags, data, ts)
	if err := writer.WritePoint(ctx, point); err != nil {
		logger.WithError(err).Debugf("failed to write data to influxdb")
		return err
	}

	logger.Debugf("write data")

	return nil
}

func NewInfluxdb2DataStorage(args ...interface{}) (DataStorage, error) {
	var logger log.FieldLogger
	var opt Influxdb2DataStorageOption

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
		opt:    &opt,
		influx: influx,
		logger: logger,
	}

	logger.Debugf("new influxdb client")

	return s, nil
}

func init() {
	registry_data_storage_factory("influxdb2", NewInfluxdb2DataStorage)
}
