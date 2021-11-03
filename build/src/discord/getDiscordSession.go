package discord

import (
	"github.com/bwmarrin/discordgo"
)

func GetDiscordSession() (*discordgo.Session, error) {
	// Create discord bot
	token := "ODk5NjUzNDY2NjAxOTQ3MTk3.YW15lQ.k9Xamjdnk2P1PQ0aGfiV8arx75Q"
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	discord.Identify.Intents = discordgo.IntentsGuildMessages
	
	return discord, nil
}

