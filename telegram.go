package scripturebot

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

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
	Id     int            `json:"message_id"`
}

type TelegramRequest struct {
	Message TelegramMessage `json:"message"`
}

type TelegramPost struct {
	Id      string `json:"chat_id"`
	Text    string `json:"text"`
	ReplyId string `json:"reply_to_message_id"`
}

type InlineButton struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

type InlineMarkup struct {
	Keyboard [][]InlineButton `json:"inline_keyboard"`
}

type TelegramInlinePost struct {
	TelegramPost
	Markup InlineMarkup `json:"reply_markup"`
}

type KeyButton struct {
	Text string `json:"text"`
}

type ReplyMarkup struct {
	Keyboard  [][]KeyButton `json:"keyboard"`
	Resize    bool          `json:"resize_keyboard"`
	Once      bool          `json:"one_time_keyboard`
	Selective bool          `json:"selective"`
}

type TelegramReplyPost struct {
	TelegramPost
	Markup ReplyMarkup `json:"reply_markup"`
}

type RemoveMarkup struct {
	Remove    bool `json:"remove_keyboard"`
	Selective bool `json:"selective"`
}

type TelegramRemovePost struct {
	TelegramPost
	Markup RemoveMarkup `json:"reply_markup`
}

func TelegramTranslate(body []byte, env *SessionData) bool {
	log.Printf("Parsing Telegram message")

	var data TelegramRequest
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal request body: %v", err)
		return false
	}

	env.User.Firstname = data.Message.Sender.Firstname
	env.User.Lastname = data.Message.Sender.Lastname
	env.User.Username = data.Message.Sender.Username
	env.User.Id = strconv.Itoa(data.Message.Sender.Id)
	env.User.Type = TYPE_TELEGRAM

	log.Printf("User: %s %s | %s : %s", env.User.Firstname, env.User.Lastname, env.User.Username, env.User.Id)

	tokens := strings.Split(data.Message.Text, " ")
	if strings.Index(tokens[0], "/") == 0 {
		env.Msg.Command = string((tokens[0])[1:])
	}
	env.Msg.Message = strings.Replace(data.Message.Text, env.Msg.Command, "", 1)
	env.Msg.Id = strconv.Itoa(data.Message.Id)

	env.Channel = strconv.Itoa(data.Message.Chat.Id)

	log.Printf("Message: %s | %s", env.Msg.Command, env.Msg.Message)

	return true
}

func PostTelegram(env *SessionData) bool {
	endpoint := "https://api.telegram.org/bot" + env.Secrets.TELEGRAM_ID + "/sendMessage"
	header := "application/json;charset=utf-8"

	var base TelegramPost
	base.Id = env.User.Id
	base.ReplyId = env.Msg.Id
	base.Text = env.Res.Message

	var data []byte
	var err error

	if env.Res.Affordances != nil {
		if len(env.Res.Affordances.Options) > 0 {
			if env.Res.Affordances.Inline {
				var buttons []InlineButton
				for i := 0; i < len(env.Res.Affordances.Options); i++ {
					buttons = append(buttons, InlineButton{env.Res.Affordances.Options[i].Text, env.Res.Affordances.Options[i].Link})
				}
				var markup InlineMarkup
				markup.Keyboard = append([][]InlineButton{}, buttons)
				var message TelegramInlinePost
				message.TelegramPost = base
				message.Markup = markup
				data, err = json.Marshal(message)
			} else {
				var buttons []KeyButton
				for i := 0; i < len(env.Res.Affordances.Options); i++ {
					buttons = append(buttons, KeyButton{env.Res.Affordances.Options[i].Text})
				}
				var markup ReplyMarkup
				markup.Keyboard = append([][]KeyButton{}, buttons)
				var message TelegramReplyPost
				message.TelegramPost = base
				message.Markup = markup
				data, err = json.Marshal(message)
			}
		} else if env.Res.Affordances.Remove {
			var message TelegramRemovePost
			message.TelegramPost = base
			message.Markup.Remove = true
			message.Markup.Selective = true
			data, err = json.Marshal(message)
		}
	} else {
		data, err = json.Marshal(base)
	}

	if err != nil {
		log.Fatalf("Error occurred during conversion to JSON: %v", err)
		return false
	}

	buffer := bytes.NewBuffer(data)
	_, err = http.Post(endpoint, header, buffer)
	if err != nil {
		log.Fatalf("Error occurred during post: %v", err)
		return false
	}

	return true
}
