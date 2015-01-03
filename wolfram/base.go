package wolfram

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/url"
)

func New(id string) *Client {
	// Make New Client Structure
	//
	var ctx Client
	ctx.appid = id

	return &ctx

}

func (c *Client) Query(input string) (*Query, error) {
	options := map[string]string{"input": input}
	query := Query{
		Options: options,
	}
	var err error
	err = c.getUnmarshal(options, &query.Results)

	return &query, err
}

func (c *Client) getUnmarshal(opts map[string]string, v interface{}) error {

	options := map[string]string{
		"appid":  c.appid,
		"format": "image,plaintext",
	}

	for key, val := range opts {
		options[key] = val
	}

	url := getWAURL(options)

	res, err := c.Client.Get(url)
	if err != nil {
		fmt.Errorf("http Get failed %v", err)
		return err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("can not read response body %v", err)
		return err
	}

	err = xml.Unmarshal(b, v)
	if err != nil {
		fmt.Errorf("unmarshal error:%v", err)
		return err
	}

	return nil
}

func flagCheck(f flag) string {

	return ""

}

func getWAURL(opts map[string]string) string {
	url, _ := url.Parse(URL_WOLFRAM_API)
	query := url.Query()

	for key, val := range opts {
		query.Add(key, val)
	}
	url.RawQuery = query.Encode()

	return url.String()
}
