package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (bot BotHandlers) MessageHandler(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
	if messageCreate.Author.ID == session.State.User.ID {
		return
	}

	if strings.Contains(messageCreate.Content, "!mygo") {
		_, err := session.ChannelMessageSend(messageCreate.ChannelID, "**燈才是最需要練習的**")
		ErrorHandle(err)
	}
}
