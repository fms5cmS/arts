
接入聊天工具并发送信息

# TelegramBot

[Telegram](https://telegram.org/)

[TelegramBot](https://core.telegram.org/bots)

[Golang Telegram Bit API](https://github.com/go-telegram-bot-api/telegram-bot-api)

- 向 @BotFather 发送消息(详情在开始聊天后会有提示)注册机器人并接收 Token
- 开启机器人消息接收程序后，获得个人 chatID 和群 chatID
  - 个人 chatID：单个用户向创建的机器人发送私信
  - 群 chatID：任意用户在群中发送`任意消息 @机器人`
    - 机器人需要有管理员权限

[简单示例](tgbot/tgbot_test.go)
