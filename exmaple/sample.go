package main

import (
	"fmt"
	"os"

	wolfram "github.com/jaisingh/go-wolfram-alpha"
)

func main() {
	id := os.Getenv("WAID")
	if id == "" {
		fmt.Println("Please define environment var WAID with your AppID")
		os.Exit(1)
	}

	c := wolfram.New(id)
	q, err := c.Get("msft")
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	fmt.Println(q)
	fmt.Println("====================")

	// example access pod
	for _, pod := range q.Pods {
		fmt.Println(pod.Title)
		for _, subpod := range pod.Subpods {
			fmt.Println("sub", subpod.Title)
		}
	}
}
