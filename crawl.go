package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/moovweb/gokogiri"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var visited = make(map[string]bool)

func main() {
	flag.Parse()

	args := flag.Args()
	fmt.Println(args)
	if len(args) < 1 {
		fmt.Println("Please specify start page")
		os.Exit(1)
	}

	queue := make(chan string)

	go func() { queue <- args[0] }()

	for uri := range queue {
		enqueue(uri, queue)
	}
}

func enqueue(uri string, queue chan string) {
	fmt.Println("fetching", uri)
	visited[uri] = true

	links, _ := findRawLinks(uri)

	for _, link := range links {
		absolute := fixUrl(link, uri)

		if uri != "" && !visited[absolute] {
			go func() { queue <- absolute }()
		}
	}
}

func fixUrl(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}

func findRawLinks(uri string) ([]string, error) {
	client := getHttpClient()

	// fetch and read a web page
	resp, err := client.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parse the web page
	doc, err := gokogiri.ParseHtml(page)
	if err != nil {
		return nil, err
	}

	// important -- don't forget to free  the resources when you're done!
	defer doc.Free()

	// perform operations on the parsed page

	var urls []string
	html := doc.Root().FirstChild()
	results, err := html.Search("//a/@href")
	if err != nil {
		return nil, err
	}

	if results != nil {
		for _, node := range results {
			urls = append(urls, node.String())
		}
	}

	return urls, nil
}

func getHttpClient() http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	return http.Client{Transport: transport}
}
