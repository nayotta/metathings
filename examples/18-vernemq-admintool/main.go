package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
	mqtt_helper "github.com/nayotta/metathings/pkg/common/mqtt"
	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
	redis_helper "github.com/nayotta/metathings/pkg/common/redis"
	metathings_plugin_vernemq_storage "github.com/nayotta/metathings/pkg/plugin/vernemq/storage"
)

var (
	redisAddr              string
	redisAddrs             []string
	redisDb                int
	redisUsername          string
	redisPassword          string
	action                 string
	vernemqClientId        string
	vernemqUsername        string
	vernemqPassword        string
	vernemqTopic           string
	vernemqTopics          []string
	vernemqSubscribeTopics []string
	vernemqPublishTopics   []string
	logLevel               string
)

func main() {
	pflag.StringVar(&redisAddr, "redis-addr", "127.0.0.1:6379", "redis addr")
	pflag.StringSliceVar(&redisAddrs, "redis-addrs", nil, "redis addrs")
	pflag.IntVar(&redisDb, "redis-db", 0, "redis db")
	pflag.StringVar(&redisUsername, "redis-username", "", "redis username")
	pflag.StringVar(&redisPassword, "redis-password", "", "redis password")
	pflag.StringVar(&action, "action", "set", "action [set, unset]")
	pflag.StringVar(&vernemqClientId, "client-id", "", "client id")
	pflag.StringVar(&vernemqUsername, "username", "", "username")
	pflag.StringVar(&vernemqPassword, "password", "", "password")
	pflag.StringVar(&vernemqTopic, "topic", "", "topic")
	pflag.StringSliceVar(&vernemqTopics, "topics", nil, "topics")
	pflag.StringSliceVar(&vernemqSubscribeTopics, "subscribe-topics", nil, "subscribe topics")
	pflag.StringSliceVar(&vernemqPublishTopics, "publish-topics", nil, "publish topics")
	pflag.StringVar(&logLevel, "log-level", "info", "log level")

	pflag.Parse()

	if redisAddr != "" {
		redisAddrs = append(redisAddrs, redisAddr)
	}

	if vernemqTopic != "" {
		vernemqTopics = append(vernemqTopics, vernemqTopic)
	}

	if vernemqSubscribeTopics == nil {
		vernemqSubscribeTopics = vernemqTopics
	}

	if vernemqPublishTopics == nil {
		vernemqPublishTopics = vernemqTopics
	}

	rsCli := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    redisAddrs,
		DB:       redisDb,
		Username: redisUsername,
		Password: redisPassword,
	})

	logger, err := log_helper.NewLogger("vernemq-admintool", logLevel)
	if err != nil {
		panic(err)
	}

	logger.WithFields(logrus.Fields{
		"redis-addr":            redisAddr,
		"redis-addrs":           redisAddrs,
		"redis-db":              redisDb,
		"redis-username":        redisUsername,
		"redis-password":        redisPassword,
		"action":                action,
		"mqtt-client-id":        vernemqClientId,
		"mqtt-username":         vernemqUsername,
		"mqtt-password":         vernemqPassword,
		"mqtt-topic":            vernemqTopic,
		"mqtt-topics":           vernemqTopics,
		"mqtt-subscribe-topics": vernemqSubscribeTopics,
		"mqtt-publish-topics":   vernemqPublishTopics,
	}).Debugf("vernemq-admintool configuration")

	storage, err := metathings_plugin_vernemq_storage.New("redis", redis_helper.WithRedisClient(rsCli), log_helper.WithLogger(logger))
	if err != nil {
		panic(err)
	}

	switch action {
	case "set":
		passwd := mqtt_helper.ParseMqttPassword(vernemqUsername, vernemqPassword)
		passhash := passwd_helper.MustParseBcrypt(passwd)
		usr := &metathings_plugin_vernemq_storage.User{
			ClientID: vernemqClientId,
			Username: vernemqUsername,
			Passhash: passhash,
		}
		if err = storage.CreateOrUpdateUser(usr); err != nil {
			panic(err)
		}
		logger.Infof("set user")
	case "unset":
		if err = storage.RemoveUser(vernemqClientId, vernemqUsername); err != nil {
			panic(err)
		}
		logger.Infof("unset user")
	default:
		panic(fmt.Sprintf("unsupported action: %s", action))
	}
}
