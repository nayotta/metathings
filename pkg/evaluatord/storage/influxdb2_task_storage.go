package metathings_evaluatord_storage

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/stretchr/objx"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

const (
	INFLUXDB2_TASK_MASUREMENT = "evaluatord.task"
)

type Influxdb2TaskStorageOption struct {
	Address string
	Token   string
	Org     string
	Bucket  string
}

type Influxdb2TaskStorage struct {
	opt    *Influxdb2TaskStorageOption
	influx influxdb2.InfluxDBClient
	logger logrus.FieldLogger
}

func (s *Influxdb2TaskStorage) get_logger() logrus.FieldLogger {
	return s.logger
}

func (s *Influxdb2TaskStorage) parse_tableresult_to_tasks(tr *influxdb2.QueryTableResult) ([]*Task, error) {
	var tsks []*Task
	tskm := map[string]*Task{}
	idtsm := map[string]map[time.Time]*TaskState{}

	for tr.Next() {
		r := tr.Record()
		rvs := r.Values()
		at := r.Time()
		id := cast.ToString(rvs["#task"])
		src := cast.ToString(rvs["#source"])
		src_typ := cast.ToString(rvs["#source_type"])
		if _, ok := tskm[id]; !ok {
			tskm[id] = &Task{
				Id: &id,
				Source: &Resource{
					Id:   &src,
					Type: &src_typ,
				},
			}
			idtsm[id] = map[time.Time]*TaskState{}
		}

		if _, ok := idtsm[id][at]; !ok {
			idtsm[id][at] = &TaskState{
				At:   &at,
				Tags: map[string]interface{}{},
			}
		}

		switch r.Field() {
		case "$state":
			v := cast.ToString(r.Value())
			idtsm[id][at].State = &v
		default:
			if len(r.Field()) > 0 && r.Field()[0] == '$' {
				idtsm[id][at].Tags[r.Field()[1:]] = r.Value()
			}
		}
	}

	if err := tr.Err(); err != nil {
		return nil, err
	}

	for id, tssm := range idtsm {
		var tss []*TaskState
		for _, sn := range tssm {
			tss = append(tss, sn)
		}
		sort.Slice(tss, func(i, j int) bool { return tss[i].At.UnixNano() < tss[j].At.UnixNano() })
		tsk := tskm[id]
		tsk.States = tss
		tsk.CurrentState = tss[len(tss)-1]
		tsk.CreatedAt = *tss[0].At
		tsk.UpdatedAt = *tss[len(tss)-1].At
		tsks = append(tsks, tsk)
	}

	return tsks, nil
}

func (s *Influxdb2TaskStorage) ListTasksBySource(ctx context.Context, src *Resource, opts ...ListTasksBySourceOption) ([]*Task, error) {
	logger := s.get_logger()

	o := make(objx.Map)
	for _, apply := range opts {
		apply(o)
	}

	query_api := s.influx.QueryApi(s.opt.Org)

	var query strings.Builder
	query.WriteString(`from(bucket: "`)
	query.WriteString(s.opt.Bucket)
	query.WriteString(`")`)
	query.WriteString(` |> range(`)
	if va, vb := o.Get("start"), o.Get("stop"); !va.IsNil() || !vb.IsNil() {
		if !va.IsNil() {
			query.WriteString(`start: `)
			query.WriteString(va.String())
		}

		if !vb.IsNil() {
			if !va.IsNil() {
				query.WriteString(`, `)
			}
			query.WriteString(`stop: `)
			query.WriteString(vb.String())
		}
	} else {
		query.WriteString(`start: -3650d`)
	}
	query.WriteString(`)`)
	query.WriteString(` |> filter(fn: (r) => r["_measurement"] == "`)
	query.WriteString(INFLUXDB2_TASK_MASUREMENT)
	query.WriteString(`")`)
	query.WriteString(` |> filter(fn: (r) => r["#source"] == "`)
	query.WriteString(*src.Id)
	query.WriteString(`" and r["#source_type"] == "`)
	query.WriteString(*src.Type)
	query.WriteString(`")`)
	query_str := query.String()
	tr, err := query_api.Query(ctx, query_str)
	if err != nil {
		logger.WithError(err).WithField("query", query_str).Debugf("failed to list tasks by source")
		return nil, err
	}

	tsks, err := s.parse_tableresult_to_tasks(tr)
	if err != nil {
		logger.WithError(err).Debugf("failed to parse table result to tasks")
		return nil, err
	}

	return tsks, nil
}

func (s *Influxdb2TaskStorage) GetTask(ctx context.Context, id string) (*Task, error) {
	logger := s.get_logger()

	query_api := s.influx.QueryApi(s.opt.Org)
	query := fmt.Sprintf(`
from(bucket: "%s")
  |> range(start: -3650d)
  |> filter(fn: (r) => r["_measurement"] == "%s")
  |> filter(fn: (r) => r["#task"] == "%s")
`,
		s.opt.Bucket,
		INFLUXDB2_TASK_MASUREMENT,
		id,
	)

	tr, err := query_api.Query(ctx, query)
	if err != nil {
		logger.WithError(err).Debugf("failed to get task")
		return nil, err
	}

	tsks, err := s.parse_tableresult_to_tasks(tr)
	if err != nil {
		logger.WithError(err).Debugf("failed to parse table result to tasks")
		return nil, err
	}

	if len(tsks) != 1 {
		err = ErrTaskNotFound
		logger.WithError(err).Debugf("task not found")
		return nil, err
	}

	return tsks[0], nil
}

func (s *Influxdb2TaskStorage) PatchTask(ctx context.Context, tsk *Task, ts *TaskState) error {
	if ts.At == nil {
		now := time.Now()
		ts.At = &now
	}

	write_api := s.influx.WriteApiBlocking(s.opt.Org, s.opt.Bucket)

	tags := map[string]interface{}{}
	for k, v := range ts.Tags {
		tags[fmt.Sprintf("$%v", k)] = v
	}
	tags["$state"] = *ts.State

	point := influxdb2.NewPoint(INFLUXDB2_TASK_MASUREMENT, map[string]string{
		"#task":        *tsk.Id,
		"#source":      *tsk.Source.Id,
		"#source_type": *tsk.Source.Type,
	}, tags, *ts.At)
	if err := write_api.WritePoint(ctx, point); err != nil {
		return err
	}

	return nil
}

func NewInfluxdb2TaskStorage(args ...interface{}) (TaskStorage, error) {
	var opt Influxdb2TaskStorageOption
	var logger logrus.FieldLogger

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger":  opt_helper.ToLogger(&logger),
		"address": opt_helper.ToString(&opt.Address),
		"token":   opt_helper.ToString(&opt.Token),
		"org":     opt_helper.ToString(&opt.Org),
		"bucket":  opt_helper.ToString(&opt.Bucket),
	}, opt_helper.SetSkip(true))(args...); err != nil {
		return nil, err
	}

	influx := influxdb2.NewClient(opt.Address, opt.Token)

	return &Influxdb2TaskStorage{
		opt:    &opt,
		influx: influx,
		logger: logger,
	}, nil
}

func init() {
	register_task_storage_factory("influxdb2", NewInfluxdb2TaskStorage)
}
