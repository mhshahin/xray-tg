package telegram

type TelegramInterface interface{}

type Telegram struct{}

func NewTelgram() TelegramInterface {
	return &Telegram{}
}
