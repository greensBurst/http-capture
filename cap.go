package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	g       bool
	p       bool
	u       string
	headers bool
	params  bool
	help    bool
)

func init() {
	flag.BoolVar(&g, "g", false, "GET REQUEST")
	flag.BoolVar(&p, "p", false, "POST REQUEST")
	flag.BoolVar(&headers, "headers", false, "ADD HEADERS")
	flag.BoolVar(&params, "params", false, "ADD PARAMS")
	flag.BoolVar(&help, "help", false, "HELP DOC")
	flag.StringVar(&u, "u", "request url", "REQUEST URL")
	flag.Parse()
}

func main() {

	var (
		req *http.Request
		res *http.Response
	)

	if u == "request url" {
		log.Fatal(errors.New("URL不能为空。"))
	}

	if g {
		req, res = GET(&u)
		defer res.Body.Close()
	} else if p {
		req, res = POST(u)
		defer res.Body.Close()
	}
	request(req)
	response(res)
	analyze(res)
}

func analyze(res *http.Response) {
	var x string
	doc, err := goquery.NewDocumentFromReader(res.Body)
	deal(err)
	for {
		fmt.Println("HTML选择器支持通过标签选择，exit结束。")
		fmt.Scanln(&x)
		if x == "exit" {
			break
		}
		doc.Find(x).Each(func(i int, selection *goquery.Selection) {
			fmt.Println(selection.Text())
			fmt.Println()
		})
	}
}

func deal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GET(url *string) (*http.Request, *http.Response) {
	getParams(url)
	req, err := http.NewRequest("GET", *url, nil)
	deal(err)
	addHeaders(req)
	res, err := http.DefaultClient.Do(req)
	deal(err)
	return req, res
}

func POST(url string) (*http.Request, *http.Response) {
	params := postParams()
	req, err := http.NewRequest("POST", url, params)
	deal(err)
	addHeaders(req)
	res, err := http.DefaultClient.Do(req)
	deal(err)
	return req, res
}

func addHeaders(req *http.Request) {
	var x string
	fmt.Println("以 key:value 格式输入要添加的消息头，exit 结束。")
	for {
		fmt.Scanln(&x)
		if x == "exit" {
			break
		}
		kv := make([]string, 2, 2)
		kv = strings.Split(x, ":")
		req.Header.Add(kv[0], kv[1])
	}
}

func getParams(url *string) {
	var (
		x string
		i int
	)
	i = 0
	fmt.Println("以 key:value 格式输入要添加的参数，exit 结束。")
	for {
		fmt.Scanln(&x)
		if x == "exit" {
			break
		}
		kv := make([]string, 2, 2)
		kv = strings.Split(x, ":")
		if i == 0 {
			*url = *url + "?" + kv[0] + "=" + kv[1]
		} else {
			*url = *url + "&" + kv[0] + "=" + kv[1]
		}
		i++
	}
}

func postParams() io.Reader {
	var (
		x   string
		i   int
		str string
	)
	i = 0
	fmt.Println("以 key:value 格式输入要添加的参数，exit 结束。")
	for {
		fmt.Scanln(&x)
		if x == "exit" {
			break
		}
		kv := make([]string, 2, 2)
		kv = strings.Split(x, ":")
		each := kv[0] + "=" + kv[1]
		if i == 0 {
			str = each
		} else {
			str = str + "&" + each
		}
		i++
	}
	params := strings.NewReader(str)
	return params
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
	for k, v := range res.Header { //v []string
		fmt.Printf("\t%v:\n", k)
		for i := 0; i < len(v); i++ {
			s := strings.Split(v[i], "; ") //i string    s []string
			for _, j := range s {
				fmt.Printf("\t\t\t\t%v\n", j)
			}
			if i != len(v)-1 {
				fmt.Println()
			}
		}
	}
}
