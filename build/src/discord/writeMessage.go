package discord

import (
	"announcements-bot/params"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func WriteNewVersionMessage(discord *discordgo.Session, discordChannel string, versionEvent *params.NewVersionEvent, packageName string) {
	message := fmt.Sprintf(":new: New DAppNode package version for %s! :up:\n\n:arrow_up: New version: **%d.%d.%d**", packageName, versionEvent.SemanticVersion[0], versionEvent.SemanticVersion[1], versionEvent.SemanticVersion[2])
	//fmt.Println(message)
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

func WriteNewRepoMessage(discord *discordgo.Session, discordChannel string, repoEvent *params.NewRepoEvent) {
	// create a custom discord message with emojis
	message := fmt.Sprintf(":new: New DAppNode package **%s**! :package:\n\n:point_right: Install it at http://my.dappnode/#/installer/%s.dnp.dappnode.eth\n:record_button: Check the transaction at https://etherscan.io/address/%s", repoEvent.Name, repoEvent.Name, repoEvent.Address)
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