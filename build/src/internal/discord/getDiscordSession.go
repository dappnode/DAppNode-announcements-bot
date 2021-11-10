package discord

import (
	"github.com/bwmarrin/discordgo"
)

func GetDiscordSession(discordToken string) (*discordgo.Session, error) {
	// Create discord bot
	discord, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		return nil, err
	}

	discord.Identify.Intents = discordgo.IntentsGuildMessages
	
	return discord, nil
}

