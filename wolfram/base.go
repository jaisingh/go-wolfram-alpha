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

func (c *Client) Get(data string, opts ...map[string]string) (*QueryResult, error) {
	//export function
	// atodeyaru
	err := c.request(data)
	return &c.QueryResult, err
}

func (c *Client) request(input string) error {
	// internal process to request wolfram-alpha.com engine.
	//
	// ***future implementation***
	// you can choice format : image, plaintext, mathematica input, and other
	// we can choice some way
	// args -> const variable like os.O_XXX
	//      -> make strucure filed {image bool,plaintext bool,mathematica_input bool}

	url := getWAURL(map[string]string{
		"appid":  c.appid,
		"input":  input,
		"format": "image,plaintext",
	})

	var query QueryResult

	//
	res, err := c.Client.Get(url)
	if err != nil {
		fmt.Errorf("http Get failed %v", err)
		return err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("can not read resoponce body %v", err)
		return err
	}

	err = xml.Unmarshal(b, &query)
	if err != nil {
		fmt.Errorf("unmarshal error:%v", err)
		return err
	}
	c.QueryResult = query

	return nil
}

func (c *Client) IsSuccessed() bool {
	//check to query-result server-side states
	return c.QueryResult.success
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
