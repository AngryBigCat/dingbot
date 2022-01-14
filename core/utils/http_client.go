package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HttpClient struct {
	http http.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

func (c *HttpClient) Get(url1 string, params1 map[string]string) (map[string]interface{}, error) {
	params := url.Values{}

	for k, v := range params1 {
		params.Set(k, v)
	}

	url2, err := url.Parse(url1)
	if err != nil {
		return nil, err
	}

	url2.RawQuery = params.Encode()

	urlPath := url2.String()

	return c.Request("GET", urlPath, nil)
}

func (c *HttpClient) Post(url string, params interface{}) (map[string]interface{}, error) {
	return c.Request("POST", url, params)
}

func (c *HttpClient) Request(method string, url string, params interface{}) (map[string]interface{}, error) {
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	postData := bytes.NewReader(jsonParams)

	req, err := http.NewRequest(method, url, postData)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	jsonData := make(map[string]interface{})

	err = json.Unmarshal(resData, &jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
