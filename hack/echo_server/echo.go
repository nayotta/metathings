package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/spf13/pflag"
)

func main() {
	var host string
	var port int

	pflag.StringVar(&host, "host", "0.0.0.0", "Host")
	pflag.IntVar(&port, "port", 8080, "Port")

	rr := http.NewServeMux()
	rr.HandleFunc("/webhook", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key := range r.Header {
			fmt.Printf("%s:\t%s\n", key, r.Header.Get(key))
		}
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var body map[string]interface{}
		err = json.Unmarshal(buf, &body)
		if err != nil {
			panic(err)
		}
		fmt.Println(body)
	}))

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}

	if err = http.Serve(lis, rr); err != nil {
		panic(err)
	}
}
