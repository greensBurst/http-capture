package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	g string
	p string
)

func init() {
	flag.StringVar(&g, "g", "get request url", "GET()请求地址")
	flag.StringVar(&p, "p", "post request url", "POST()请求地址")
	flag.Parse()
}

func main() {

	var (
		req *http.Request
		res *http.Response
	)
	if g != "get request url" {
		req, res = newreq("GET", g, nil)
		request(req)
		response(res)
	} else if p != "post request url" {
		form := make(map[string]string, 0)
		fmt.Println("input key:value in one line,press over to break.")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == "over" {
				break
			}
			kv := make([]string, 2, 2)
			kv = strings.Split(scanner.Text(), ":")
			form[kv[0]] = kv[1]
		}

		info := makeParameter(form)
		req, res = newreq("POST", p, info)
		request(req)
		response(res)
	}
	defer res.Body.Close()
	fmt.Println("You can enter a html element that you want.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dom(scanner.Text(), res.Body)
}

func dom(label string, Body io.Reader) {
	doc, _ := goquery.NewDocumentFromReader(Body)
	doc.Find(label).Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
		fmt.Println()
	})
}

func makeParameter(form map[string]string) io.Reader {
	var str string
	cnt := 0
	for k, v := range form {
		each := k + "=" + v
		if cnt == 0 {
			str = each
		} else {
			str = str + "&" + each
		}
		cnt++
	}
	info := strings.NewReader(str)
	return info
}

func deal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func newreq(method string, url string, parameter io.Reader) (*http.Request, *http.Response) {
	req, err := http.NewRequest(method, url, parameter)
	deal(err)
	if method == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	res, err := http.DefaultClient.Do(req)
	deal(err)
	// defer res.Body.Close()
	return req, res
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
