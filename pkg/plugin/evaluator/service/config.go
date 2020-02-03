package main

type Config = Data
type config = data

func ConfigFromMap(raw map[string]interface{}) (Config, error) {
	return &config{raw: raw}, nil
}

func ConfigFromBytes(buf []byte) (Config, error) {
	return &config{buf: buf}, nil
}
