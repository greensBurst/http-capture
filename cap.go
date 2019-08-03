package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var (
	g string
)

func init() {
	flag.StringVar(&g, "g", "get request url", "GET()请求地址")
}

func main() {
	flag.Parse()

	if g != "get request url" {
		res, err := http.Get(g)
		deal(err)
		defer res.Body.Close()
		response(res)
		request(res.Request)
	}
}

func deal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func request(req *http.Request) {
	fmt.Println("\nrequest Line:")
	fmt.Printf("\tMethod:\t\t%v\n\tURL:\t\t%v\n\tProtocol:\t%v\n", req.Method, req.URL, req.Proto)

	fmt.Println("\nrequest Header:")
	for k, v := range req.Header {
		fmt.Printf("\t%v:\n", k)
		for i := 0; i < len(v); i++ {
			s := strings.Split(v[i], "; ")
			for _, j := range s {
				fmt.Printf("\t\t\t\t%v\n", j)
			}
			if i != len(v)-1 {
				fmt.Println()
			}
		}
	}
}

func response(res *http.Response) {
	fmt.Println("\nresponse Line:")
	fmt.Printf("\tMethod:\t\t%v\n\tStatus:\t\t%v\n\tProtocol:\t%v\n", res.Request.Method, res.Status, res.Proto)

	fmt.Println("\nresponse Header:")
	for k, v := range res.Header {
		fmt.Printf("\t%v:\n", k)
		for i := 0; i < len(v); i++ {
			s := strings.Split(v[i], "; ")
			for _, j := range s {
				fmt.Printf("\t\t\t\t%v\n", j)
			}
			if i != len(v)-1 {
				fmt.Println()
			}
		}
	}
}
