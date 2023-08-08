package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"io"
	"os"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
	metathings_module_soda_sdk "github.com/nayotta/metathings/sdk/module/soda"
)

func main() {
	var url string
	var source, destination string

	flag.StringVar(&url, "url", "http://127.0.0.1:8001", "")
	flag.StringVar(&source, "source", "", "")
	flag.StringVar(&destination, "destination", "", "")
	flag.Parse()

	logger, err := log_helper.NewLogger("sdk", "trace")
	if err != nil {
		panic(err)
	}

	cli, err := metathings_module_soda_sdk.NewSodaClient(
		metathings_module_soda_sdk.WithLogger(logger.WithFields(nil)),
		metathings_module_soda_sdk.WithURL(url),
	)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(source)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fs, err := os.Stat(source)
	if err != nil {
		panic(err)
	}

	h := sha1.New()
	bufio.NewReader(f).WriteTo(h)
	buf := h.Sum(nil)
	sha1sum := hex.EncodeToString(buf)

	f.Seek(0, io.SeekStart)

	err = cli.PutObjectStreaming(destination, f, fs.Size(), metathings_module_soda_sdk.PutObjectStreamingOption{Sha1sum: sha1sum})
	if err != nil {
		panic(err)
	}
}
