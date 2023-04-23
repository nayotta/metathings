package test_helper

import (
	"os"
	"strings"

	"github.com/spf13/cast"
)

func GetenvWithDefault(k, d string) string {
	v := os.Getenv(k)
	if v == "" {
		v = d
	}
	return v
}

func GetTestGormDriver() string {
	return GetenvWithDefault("MTT_GORM_DRIVER", "sqlite3")
}

func GetTestGormUri() string {
	return GetenvWithDefault("MTT_GORM_URI", ":memory:")
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

func GetTestInfluxdb2Address() string {
	return GetenvWithDefault("MTT_INFLUXDB2_ADDRESS", "http://localhost:9999")
}

func GetTestInfluxdb2Token() string {
	return GetenvWithDefault("MTT_INFLUXDB2_TOKEN", "")
}

func GetTestInfluxdb2Org() string {
	return GetenvWithDefault("MTT_INFLUXDB2_ORG", "")
}

func GetTestInfluxdb2Bucket() string {
	return GetenvWithDefault("MTT_INFLUXDB2_BUCKET", "")
}

func GetTestMinioEndpoint() string {
	return GetenvWithDefault("MTT_MINIO_ENDPOINT", "127.0.0.1:9000")
}

func GetTestMinioID() string {
	return GetenvWithDefault("MTT_MINIO_ID", "minioadmin")
}

func GetTestMinioSecret() string {
	return GetenvWithDefault("MTT_MINIO_SECRET", "minioadmin")
}

func GetTestMinioToken() string {
	return GetenvWithDefault("MTT_MINIO_TOKEN", "")
}

func GetTestMinioSecure() bool {
	return cast.ToBool(GetenvWithDefault("MTT_MINIO_SECURE", "false"))
}

func GetTestMinioBucket() string {
	return GetenvWithDefault("MTT_MINIO_BUCKET", "test")
}

func GetTestMinioReadBufferSize() int {
	return cast.ToInt(GetenvWithDefault("MTT_MINIO_READ_BUFFER_SIZE", "4194304")) // 4MiB
}

func GetTestMinioWriteBufferSize() int {
	return cast.ToInt(GetenvWithDefault("MTT_MINIO_WRITE_BUFFER_SIZE", "4194304")) // 4MiB
}
