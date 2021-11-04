package main

import (
	"announcements-bot/discord"
	"announcements-bot/eth"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var gethRpc , discordToken , discordChannel string


func init() {
	fmt.Println("Setting up environment variables")

	environment := os.Getenv("GO_ENV")
	if environment == "development" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
		}
		gethRpc = os.Getenv("GETH_RPC")
		discordToken = os.Getenv("DISCORD_TOKEN")
		discordChannel = os.Getenv("ANNOUNCEMENTS_CHANNEL_ID")
	} else if environment == "production" {

	} else {
		panic("Environment not set")
	}

	// throw error if env vars are not set
	if gethRpc == "" {
		panic("gethRpc not set")
	}
	if discordToken == "" {
		panic("discordToken not set")
	}
	if discordChannel == "" {
		panic("discordChannel not set")
	}
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

	// Get repossitories
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
	
	// Start go rutine eth suscription to "NewVersion" event of all repos addresses
	fmt.Println("Go rutine NewVersion event subscription...")
	go eth.SubscribeNewVersion(ethClient, discord, discordChannel, repos)

	// Start go rutine eth suscription to "NewRepo" event of Registry 
	fmt.Println("Go rutine NewRepo event subscription...")
	go eth.SubscribeNewRepo(ethClient, discord, discordChannel)

	fmt.Println("Application successfully started")
	// Wait foerever
	select{}
}
