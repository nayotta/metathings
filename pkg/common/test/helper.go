package test_helper

import "os"

func GetenvWithDefault(k, d string) string {
	v := os.Getenv(k)
	if v == "" {
		v = d
	}
	return d
}

func GetTestMongoUri() string {
	return GetenvWithDefault("MTT_MONGO_URI", "127.0.0.1:27017")
}

func GetTestMongoDatabase() string {
	return GetenvWithDefault("MTT_MONGO_DATABASE", "mttest")
}

func GetTestMongoCollection() string {
	return GetenvWithDefault("MTT_MONGO_COLLECTION", "mttest")
}
