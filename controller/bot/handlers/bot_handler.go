package handlers

import (
	"errors"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

type BotHandlers struct {
	BotConfig DiscordBotConfig
}

func ErrorHandle(err error) {
	if err != nil {
		log.Println("something is error, pls contact support!")
		log.Fatal(err)
		return
	}
}

type DiscordBotConfig struct {
	Token string
}

func GetBotToken() string {
	botToken := os.Getenv("MYGO_BOT_TOKEN")
	if botToken == "" {
		ErrorHandle(errors.New("no bot token found"))
	}

	log.Println("Get bot Token: " + string(botToken))
	return string(botToken)
}

func GetConfigWithToken(botToken string) DiscordBotConfig {
	return DiscordBotConfig{Token: botToken}
}

func SendMessage(session *discordgo.Session, channelID string, message string) {
	_, err := session.ChannelMessageSend(channelID, message)
	ErrorHandle(err)
}
