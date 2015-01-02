package main

import (
	"fmt"

	"os"

	"github.com/jaisingh/go-wolfram-alpha/wolfram"
)

func main() {
	id := os.Getenv("WAID")
	if id == "" {
		fmt.Println("Please define environment var WAID with your AppID")
		os.Exit(1)
	}

	c := wolfram.New(id)

	/* proxyUrl, _ := url.Parse("http://localhost:8080")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy:           http.ProxyURL(proxyUrl),
	}
	httpClient := http.Client{Transport: tr}
	c.Client = httpClient
	*/

	q, err := c.Get("msft")
	if err != nil {
		fmt.Print("Error:", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", q)
	fmt.Println("====================")

	// example access pod
	for _, pod := range q.Pods {
		fmt.Println(pod.Title)
		for _, subpod := range pod.Subpods {
			fmt.Println("sub", subpod.Title)
		}
	}
}
