package bot

import (
	"log"
	"mygo/controller/bot/handlers"

	"github.com/bwmarrin/discordgo"
)

var goBot *discordgo.Session

func Start() {
	botConfig := handlers.GetConfigWithToken(handlers.GetBotToken())
	botHandlers := handlers.BotHandlers{BotConfig: botConfig}
	goBot, _ = discordgo.New("Bot " + botHandlers.BotConfig.Token)

	goBot.AddHandler(botHandlers.MessageHandler)
	// goBot.AddHandler(botHandlers.VoiceChannelHandler)
	// goBot.AddHandler(botHandlers.VoiceDelayHandler)

	err := goBot.Open()
	handlers.ErrorHandle(err)
	log.Println("Start Discord Bot")
}
