package utils

import (
	"ding/core"
	"fmt"
	"testing"
)

func TestJuheGet(t *testing.T) {
	data, err := NewHttpClient().JuheGet(core.API_EXCHANGE_URL, map[string]string{
		"key": core.API_EXCHANGE_KEY,
	})

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(data)
}
