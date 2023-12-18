package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"testing"
)

var botToken = "your bot token"
var chatID int64 = 0000000
var groupChatID int64 = -111111111

// 处理机器人收到的命令
// 可用于获取个人 chatID、群 chatID
func TestTGBotReceive(t *testing.T) {
	botAPI, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("New tg bot api failed: %s", err)
	}
	botAPI.Debug = true
	log.Printf("Authorized on account %s", botAPI.Self.UserName)

	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	u := tgbotapi.NewUpdate(0)
	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	u.Timeout = 30

	// 获取Bot的更新通道
	updates := botAPI.GetUpdatesChan(u)

	// 处理收到的更新
	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		// 判断处理各种命令
		message := ""
		if update.Message.IsCommand() {
			command := update.Message.Command()
			switch command {
			case "start":
				message = "Hello! I am your Telegram Bot."
			case "custom":
				message = "This is a custom command."
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			botAPI.Send(msg)
		}
	}
}

func TestTGBotSend(t *testing.T) {
	botAPI, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("New tg bot api failed: %s", err)
	}
	botAPI.Debug = true
	log.Printf("Authorized on account %s", botAPI.Self.UserName)
	message := "process run error"
	msg := tgbotapi.NewMessage(groupChatID, message)
	if _, err = botAPI.Send(msg); err != nil {
		log.Fatalf("tg bot send msg failed: %s", err)
	}
}
