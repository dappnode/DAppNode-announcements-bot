package main

import (
	"fmt"

	"github.com/pablomendezroyo/DAppNode-announcements-bot/internal/discord"
	"github.com/pablomendezroyo/DAppNode-announcements-bot/internal/env"
	"github.com/pablomendezroyo/DAppNode-announcements-bot/internal/eth"
)

var gethRpc , discordToken , discordChannel string


func init() {
	gethRpc, discordToken , discordChannel = env.LoadEnv()
} 
  
func main() {
	fmt.Println("=== Starting application ===")

	// Connect to ETH node
	fmt.Println("Connecting to Ethereum node...")
	ethClient, err := eth.GetEthClient(gethRpc)
	if err != nil {
		err := fmt.Errorf("unable to connect to %s. %s", gethRpc, err)
		fmt.Println(err)
	}

	// Get repositories
	fmt.Println("Getting DAppNode packages...")
	repos, err := eth.GetRepos(ethClient)
	if err != nil {
		err := fmt.Errorf("unable to get repositories: %w", err)
		fmt.Println(err)
	}

	// Get discord session
	fmt.Println("Getting Discord session...")
	discord, err := discord.GetDiscordSession(discordToken)
	if err != nil {
		err := fmt.Errorf("unable to open discord session: %w", err)
		fmt.Println(err)
	}

	eth.WriteNewRepoMessage(discord, discordChannel, &repos[0])
	
	// Start go rutine eth suscription to "NewVersion" event of all repos addresses
/* 	fmt.Println("Go rutine NewVersion event subscription...")
	go eth.SubscribeNewVersion(ethClient, discord, discordChannel, repos) */

	// Start go rutine eth suscription to "NewRepo" event of Registry 
/* 	fmt.Println("Go rutine NewRepo event subscription...")
	go eth.SubscribeNewRepo(ethClient, discord, discordChannel) */

	fmt.Println("Application successfully started")
	// Wait foerever
	//select{}
}
