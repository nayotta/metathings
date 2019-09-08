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

func GetTestRedisAddr() string {
	return GetenvWithDefault("MTT_REDIS_ADDR", "127.0.0.1:6379")
}

func GetTestRedisPassword() string {
	return GetenvWithDefault("MTT_REDIS_PASSWORD", "")
}

func GetTestRedisDB() string {
	return GetenvWithDefault("MTT_REDIS_DB", "0")
}

func GetTestPolicydAddress() string {
	return GetenvWithDefault("MTT_POLICYD_ADDRESS", "localhost:21733")
}

func GetTestPolicydDriverName() string {
	return GetenvWithDefault("MTT_POLICYD_DRIVER_NAME", "postgres")
}

func GetTestPolicydConnectString() string {
	return GetenvWithDefault("MTT_POLICYD_CONNECT_STRING", "host=127.0.0.1 port=5432 user=postgres password=postgres sslmode=disable")
}
