package utils

import (
	"encoding/json"
	"errors"
)

func (c *HttpClient) JuheGet(url string, params map[string]string) (string, error) {
	res, err := c.Get(url, params)
	if err != nil {
		return "", err
	}

	if int(res["error_code"].(float64)) != 0 {
		return "", errors.New(res["reason"].(string))
	}

	data, err := json.Marshal(res["result"])
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (c *HttpClient) JuhePost(url string, params map[string]string) (string, error) {
	res, err := c.Post(url, params)
	if err != nil {
		return "", err
	}

	if res["error_code"] != 0 {
		return "", errors.New(res["reason"].(string))
	}

	data, err := json.Marshal(res["result"])
	if err != nil {
		return "", err
	}

	return string(data), nil
}
