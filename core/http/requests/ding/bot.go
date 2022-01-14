package ding

type BotRequestBody struct {
	ConversationId string `json:"conversationId"`
	AtUsers        []struct {
		DingtalkId string `json:"dingtalkId"`
		StaffId    string `json:"staffId"`
	}
	ChatbotCorpId             string `json:"chatbotCorpId"`
	ChatbotUserId             string `json:"chatbotUserId"`
	MsgId                     string `json:"msgId"`
	SenderNick                string `json:"senderNick"`
	IsAdmin                   bool   `json:"isAdmin"`
	SenderStaffId             string `json:"senderStaffId"`
	SessionWebhookExpiredTime int64  `json:"sessionWebhookExpiredTime"`
	CreateAt                  int64  `json:"createAt"`
	SenderCorpId              string `json:"senderCorpId"`
	ConversationType          string `json:"conversationType"`
	SenderId                  string `json:"senderId"`
	ConversationTitle         string `json:"conversationTitle"`
	IsInAtList                bool   `json:"isInAtList"`
	SessionWebhook            string `json:"sessionWebhook"`
	Text                      struct {
		Content string `json:"content"`
	}
	Msgtype string `json:"msgtype"`
}

type BotRequestHeader struct {
	ContentType string `json:"contentType"`
	Timestamp   string `json:"timestamp"`
	Sign        string `json:"sign"`
}

type BotRequest struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}
