package main

import (
	"announcements-bot/env"
	"announcements-bot/eth"
	"announcements-bot/repository"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/ethclient"
)

var gethRpc, discordToken, discordChannel string

func init() {
	env.LoadEnv()
	// Load env vars
	gethRpc = os.Getenv("GETH_RPC")
	discordToken = os.Getenv("DISCORD_TOKEN")
	discordChannel = os.Getenv("ANNOUNCEMENTS_CHANNEL_ID")

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
	ethClient, err := getEthClient(gethRpc)
	if err != nil {
		err := fmt.Errorf("unable to connect to %s. %s", gethRpc, err)
		fmt.Println(err)
	}

	// Get repositories
	fmt.Println("Getting DAppNode packages...")
	repos, err := repository.GetRepos(ethClient)
	if err != nil {
		err := fmt.Errorf("unable to get repositories: %w", err)
		fmt.Println(err)
	}

	fmt.Println(repos)

	// Get discord session
	fmt.Println("Getting Discord session...")
	dc, err := getDiscordSession(discordToken)
	if err != nil {
		err := fmt.Errorf("unable to open discord session: %w", err)
		fmt.Println(err)
	}

	// Start go rutine eth suscription to "NewVersion" event of all repos addresses
	fmt.Println("Go rutine NewVersion event subscription...")
	go eth.SubscribeNewVersion(ethClient, dc, discordChannel, repos) 

	// Start go rutine eth suscription to "NewRepo" event of Registry 
 	fmt.Println("Go rutine NewRepo event subscription...")
	go eth.SubscribeNewRepo(ethClient, dc, discordChannel) 

	fmt.Println("Application successfully started")
	// Wait foerever
	select{}
}

func getEthClient(rpc string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getDiscordSession(discordToken string) (*discordgo.Session, error) {
	// Create discord bot
	discord, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		return nil, err
	}

	discord.Identify.Intents = discordgo.IntentsGuildMessages
	
	return discord, nil
}
