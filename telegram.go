package scripturebot

type TelegramSender struct {
	Id        int    `json:"id"`
	Bot       bool   `json:"is_bot"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Username  string `json:"username"`
	Language  string `json:"langauge_code"`
}

type TelegramChat struct {
	Id        int    `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type TelegramMessage struct {
	Sender TelegramSender `json:"from"`
	Chat   TelegramChat   `json:"chat"`
	Text   string         `json:"text"`
}

type TelegramRequest struct {
	Message TelegramMessage `json:"message"`
}
