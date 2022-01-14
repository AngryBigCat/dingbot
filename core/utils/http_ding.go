package utils

import (
	"ding/core"
	"errors"
	"log"
)

func (c *HttpClient) DingBotPost(params map[string]interface{}) error {
	url := getDingBotApi()

	res, err := c.Post(url, params)
	if err != nil {
		return err
	}

	log.Println("ding post: " + res["errmsg"].(string))

	if int64(res["errcode"].(float64)) != 0 {
		return errors.New(res["errmsg"].(string))
	}

	return nil
}

func getDingBotApi() string {
	return core.DING_API + "?access_token=" + core.DING_ACCESS_TOKEN
}
