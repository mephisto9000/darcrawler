package main

import (
  "crypto/tls"   
  "flag"         
  "fmt"
  "github.com/jackdanger/collectlinks"
  "net/http"
  "os" 
)

func main() {
  flag.Parse()

  args := flag.Args()
  fmt.Println(args)
  if len(args) < 1 {
    fmt.Println("Please specify start page")
    os.Exit(1)
  }


	queue := make(chan string)

	go func() {

		queue <- args[0]
	}()

	for uri := range queue {

	enqueue(uri, queue)
	}
	}

	func enqueue(uri string, queue chan string) {
		fmt.Println("fetching", uri)

		tlsConfig := &tls.Config{
			InsertSkipVerify: true,
		}

		transport := &http.Transport{
			TLSClientConfig: tlsConfig,
		}

		client := http.Client{Transport: transport}
		resp, err := client.Get(uri)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		links := collectlinks.All(resp.Body)

		for _, link := range links {
			go func() {queue <- link} ()
		}
	}

 