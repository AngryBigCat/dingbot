package controllers

import (
	"ding/core"
	"ding/core/http/requests/ding"
	"log"
	"strings"

	"ding/core/utils"

	"github.com/gin-gonic/gin"
)

type DingController struct {
	BaseController
}

func (cc DingController) Bot(c *gin.Context) {

	header := ding.BotRequestHeader{}
	message := ding.BotRequestBody{}

	c.BindHeader(&header)
	c.BindJSON(&message)

	contentData := message.Text.Content

	log.Println("input content: " + contentData)

	cmd, param1 := parseCommand(contentData)

	content, err := execCommand(cmd, param1...)
	if err != nil {
		content = err.Error()
	}

	textParam := map[string]string{
		"content": content,
	}
	params := map[string]interface{}{
		"msgtype": "text",
		"text":    textParam,
	}

	utils.NewHttpClient().DingBotPost(params)

	c.JSON(200, gin.H{
		"msg": "ok",
	})
}

func parseCommand(cmd string) (string, []string) {
	var cmdArr1 []string

	cmdArr := strings.Split(cmd, " ")

	for _, v := range cmdArr {
		if len(v) != 0 {
			cmdArr1 = append(cmdArr1, v)
		}
	}

	return cmdArr1[0], cmdArr1[1:]
}

func execCommand(cmd string, params ...string) (string, error) {
	log.Println("exec content: " + cmd)
	switch cmd {
	case "/help":
		return helpCmd()
	case "/look":
		return lookCmd(params[0])
	case "/exchange":
		return exchangeCmd()
	case "/ip":
		return ipCmd(params[0])
	case "/phone":
		return phoneCmd(params[0])
	case "/news":
		return newsCmd()
	default:
		return helpCmd()
	}
}

func helpCmd() (string, error) {
	content := `支持的命令：
/help 			帮助信息
/look {name} 	淘宝信誉查号
/exchange		当前汇率
/ip {ip}		查询IP归属
/phone {phone}	查询IP归属
/news			查询IP归属
`
	return content, nil
}

func lookCmd(name string) (string, error) {
	return "", nil
}

func exchangeCmd() (string, error) {
	return utils.NewHttpClient().JuheGet(core.API_EXCHANGE_URL, map[string]string{
		"key": core.API_EXCHANGE_KEY,
	})
}

func newsCmd() (string, error) {
	return utils.NewHttpClient().JuheGet(core.API_NEWS_URL, map[string]string{
		"key": core.API_NEWS_KEY,
	})
}

func ipCmd(ip string) (string, error) {
	return utils.NewHttpClient().JuheGet(core.API_IP_URL, map[string]string{
		"key": core.API_IP_KEY,
		"ip":  ip,
	})
}

func phoneCmd(phone string) (string, error) {
	return utils.NewHttpClient().JuheGet(core.API_PHONE_URL, map[string]string{
		"key":   core.API_PHONE_KEY,
		"phone": phone,
	})
}
