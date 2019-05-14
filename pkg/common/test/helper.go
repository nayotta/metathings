package test_helper

import (
	"os"
	"strings"
)

func GetenvWithDefault(k, d string) string {
	v := os.Getenv(k)
	if v == "" {
		v = d
	}
	return v
}

func GetTestMongoUri() string {
	return GetenvWithDefault("MTT_MONGO_URI", "mongodb://127.0.0.1:27017")
}

func GetTestMongoDatabase() string {
	return GetenvWithDefault("MTT_MONGO_DATABASE", "test")
}

func GetTestMongoCollection() string {
	return GetenvWithDefault("MTT_MONGO_COLLECTION", "test")
}

func GetTestKafkaBrokers() []string {
	brokers := GetenvWithDefault("MTT_KAFKA_BROKERS", "127.0.0.1:9092")
	return strings.Split(brokers, ",")
}

func GetTestPolicydAddress() string {
	return GetenvWithDefault("MTT_POLICYD_ADDRESS", "localhost:21733")
}
