package utils

import (
	"testing"
)

func TestDingBotPost(t *testing.T) {
	text := map[string]string{
		"content": "zxczxc",
	}

	params := map[string]interface{}{
		"msgtype": "text",
		"text":    text,
	}

	err := NewHttpClient().DingBotPost(params)
	if err != nil {
		t.Error(err.Error())
	}
}
