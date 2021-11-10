package discord

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func WriteNewVersionMessage(discord *discordgo.Session, discordChannel string, versionEvent *NewVersionEvent) {
	message := fmt.Sprintf("New DAppNode package version!\nVersionId: %d\nSemantic version: %d", versionEvent.versionId, versionEvent.semanticVersion)
	err := discord.Open()
	if err != nil {
		err := fmt.Errorf("unable to open discord session: %w", err)
		fmt.Println(err)
		return
	}

	_, err = discord.ChannelMessageSend(discordChannel, message)
	if err != nil {
		err := fmt.Errorf("unable to send NewVersion message: %w", err)
		log.Fatal(err)
	}
	discord.Close()
}

func WriteNewRepoMessage(discord *discordgo.Session, discordChannel string, repoEvent *NewRepoEvent) {
	message := fmt.Sprintf("New DAppNode package %s!", repoEvent.name)
	err := discord.Open()
	if err != nil {
		err := fmt.Errorf("unable to open discord session: %w", err)
		log.Fatal(err)
		return
	}

	_, err = discord.ChannelMessageSend(discordChannel, message)
	if err != nil {
		err := fmt.Errorf("unable to send NewRepo message: %w", err)
		log.Fatal(err)
	}
	discord.Close()
}