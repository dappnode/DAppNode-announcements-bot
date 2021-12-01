package discord

import (
	"announcements-bot/params"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func WriteNewVersionMessage(discord *discordgo.Session, discordChannel string, versionEvent *params.NewVersionEvent, packageName string) {
	message := fmt.Sprintf(":new: New DAppNode package version for **%s**! :up:\n\n:arrow_up: New version: **%d.%d.%d**\n\n:point_right: Update it here **http://my.dappnode/#/installer/%s.dnp.dappnode.eth**", packageName, versionEvent.SemanticVersion[0], versionEvent.SemanticVersion[1], versionEvent.SemanticVersion[2], packageName)
	err := discord.Open()
	if err != nil {
		fmt.Printf(params.WarnLog + "unable to open discord session: %w\n", err)
		return
	}

	_, err = discord.ChannelMessageSend(discordChannel, message)
	if err != nil {
		fmt.Printf(params.WarnLog + "unable to send NewVersion message: %w\n", err)
	}
	discord.Close()
}

func WriteNewRepoMessage(discord *discordgo.Session, discordChannel string, repoEvent *params.NewRepoEvent) {
	// create a custom discord message with emojis
	message := fmt.Sprintf(":new: @everyone New DAppNode package **%s**! :package:\n\n:point_right: Install it here http://my.dappnode/#/installer/%s.dnp.dappnode.eth\n\n:record_button: Check the transaction at https://etherscan.io/address/%s", repoEvent.Name, repoEvent.Name, repoEvent.Address)
	err := discord.Open()
	if err != nil {
		fmt.Printf(params.WarnLog + "unable to open discord session: %w\n", err)
		return
	}

	_, err = discord.ChannelMessageSend(discordChannel, message)
	if err != nil {
		fmt.Printf(params.WarnLog + "unable to send NewPackage message: %w\n", err)
	}
	discord.Close()
}