package token_helper

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	const_helper "github.com/nayotta/metathings/pkg/common/constant"
	identityd2_contrib "github.com/nayotta/metathings/pkg/identityd2/contrib"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

var (
	DEFAULT_REFRESH_PERIOD  = 12 * time.Hour
	MIN_REFRESH_PERIOD      = 15 * time.Second
	MAX_REFRESH_PERIOD      = 8 * time.Minute
	MAX_RETRY_REFRESH_COUNT = int64(12)
)

type Tokener interface {
	GetToken() string
}

type tokener struct {
	mtx               *sync.Mutex
	cli_fty           *client_helper.ClientFactory
	nonexpire         bool
	credential_domain string
	credential_id     string
	credential_secret string
	credential_token  string
	logger            logrus.FieldLogger

	current_refresh_period      time.Duration
	default_refresh_period      time.Duration
	min_refresh_period          time.Duration
	max_refresh_period          time.Duration
	max_retry_refresh_count     int64
	current_retry_refresh_count int64
}

func (self *tokener) GetToken() string {
	self.mtx.Lock()
	defer self.mtx.Unlock()

	return "Bearer " + self.credential_token
}

func (self *tokener) issueToken() error {
	self.mtx.Lock()
	defer self.mtx.Unlock()

	logger := self.logger

	cli, cfn, err := self.cli_fty.NewIdentityd2ServiceClient()
	if err != nil {
		logger.WithError(err).Debugf("failed to new identityd2 service client")
		return err
	}
	defer cfn()

	itbc_req := identityd2_contrib.NewIssueTokenByCredentialRequest(const_helper.DEFAULT_DOMAIN, self.credential_id, self.credential_secret)
	itbc_res, err := cli.IssueTokenByCredential(context.Background(), itbc_req)
	if err != nil {
		logger.WithError(err).Debugf("failed to issue token by credential")
		return err
	}
	txt := itbc_res.Token.Text

	if self.nonexpire {
		itbt_req := identityd2_contrib.NewIssueTokenByTokenRequest(self.credential_domain, txt)
		itbt_res, err := cli.IssueTokenByToken(context.Background(), itbt_req)
		if err != nil {
			logger.WithError(err).Debugf("failed to issue token by token")
			return err
		}
		txt = itbt_res.Token.Text
	}

	self.credential_token = txt
	self.logger.WithField("token", txt[:8]).Debugf("issue token")

	return nil
}

func (self *tokener) refreshTokenLoop() {
	for {
		if err := self.refreshTokenOnce(); err != nil {
			if self.current_refresh_period == self.default_refresh_period {
				self.current_refresh_period = self.min_refresh_period
			} else {
				self.current_refresh_period = time.Duration(math.Min(float64(self.current_refresh_period*2), float64(self.max_refresh_period)))
			}
			self.current_retry_refresh_count += 1
		} else {
			self.current_refresh_period = self.default_refresh_period
			self.current_retry_refresh_count = 0
		}

		if self.current_retry_refresh_count > self.max_retry_refresh_count {
			if err := self.issueToken(); err == nil {
				self.current_refresh_period = self.default_refresh_period
				self.current_retry_refresh_count = 0
			}
		}

		time.Sleep(self.current_refresh_period)
	}
}

func (self *tokener) refreshTokenOnce() error {
	self.mtx.Lock()
	defer self.mtx.Unlock()

	cli, cfn, err := self.cli_fty.NewIdentityd2ServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	ct_req := &pb.CheckTokenRequest{
		Token: &pb.OpToken{
			Text: &wrappers.StringValue{Value: self.credential_token},
		},
	}

	if _, err = cli.CheckToken(context.TODO(), ct_req); err != nil {
		return err
	}

	return nil
}

func NewTokener(cli_fty *client_helper.ClientFactory, credential_domain, credential_id, credential_secret string, logger log.FieldLogger) (Tokener, error) {
	tknr := &tokener{
		mtx:                         new(sync.Mutex),
		logger:                      logger,
		cli_fty:                     cli_fty,
		nonexpire:                   false,
		credential_domain:           credential_domain,
		credential_id:               credential_id,
		credential_secret:           credential_secret,
		current_refresh_period:      DEFAULT_REFRESH_PERIOD,
		default_refresh_period:      DEFAULT_REFRESH_PERIOD,
		min_refresh_period:          MIN_REFRESH_PERIOD,
		max_refresh_period:          MAX_REFRESH_PERIOD,
		max_retry_refresh_count:     MAX_RETRY_REFRESH_COUNT,
		current_retry_refresh_count: 0,
	}

	if err := tknr.issueToken(); err != nil {
		return nil, err
	}

	return tknr, nil
}

func NewNoExpireTokener(cli_fty *client_helper.ClientFactory, credential_domain, credential_id, credential_secret string, logger logrus.FieldLogger) (Tokener, error) {
	tknr := &tokener{
		mtx:                         new(sync.Mutex),
		logger:                      logger,
		cli_fty:                     cli_fty,
		nonexpire:                   true,
		credential_domain:           credential_domain,
		credential_id:               credential_id,
		credential_secret:           credential_secret,
		current_refresh_period:      DEFAULT_REFRESH_PERIOD,
		default_refresh_period:      DEFAULT_REFRESH_PERIOD,
		min_refresh_period:          MIN_REFRESH_PERIOD,
		max_refresh_period:          MAX_REFRESH_PERIOD,
		max_retry_refresh_count:     MAX_RETRY_REFRESH_COUNT,
		current_retry_refresh_count: 0,
	}

	if err := tknr.issueToken(); err != nil {
		return nil, err
	}

	go tknr.refreshTokenLoop()

	return tknr, nil
}
