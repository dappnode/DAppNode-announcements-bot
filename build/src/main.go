package main

import (
	"announcements-bot/env"
	"announcements-bot/eth"
	"announcements-bot/params"
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
		panic(params.ErrorLog + "gethRpc not set")
	}
	if discordToken == "" {
		panic(params.ErrorLog + "discordToken not set")
	}
	if discordChannel == "" {
		panic(params.ErrorLog + "discordChannel not set")
	}
} 
  
func main() {
	fmt.Println(params.InfoLog + "============================")
	fmt.Println(params.InfoLog + "=== Starting application ===")
	fmt.Println(params.InfoLog + "============================")					  

	// Connect to ETH node
	fmt.Println(params.InfoLog + "Connecting to Ethereum node...")
	ethClient, err := getEthClient(gethRpc)
	if err != nil {
		err := fmt.Errorf("unable to connect to %s. %s", gethRpc, err)
		panic(params.ErrorLog + err.Error())
	}

	// Get repositories
	fmt.Println(params.InfoLog + "Getting DAppNode packages...")
	repos, err := repository.GetRepos(ethClient)
	if err != nil {
		err := fmt.Errorf("unable to get repositories: %w", err)
		panic(params.ErrorLog + err.Error())
	}

	// Print repositories to be subscribed to
	names := repository.GetNames(repos)
	fmt.Printf(params.InfoLog + "Repositories to be subscribed to: %s\n", names)

	// Get discord session
	fmt.Println(params.InfoLog + "Getting Discord session...")
	dc, err := getDiscordSession(discordToken)
	if err != nil {
		err := fmt.Errorf("unable to open discord session: %w", err)
		panic(params.ErrorLog + err.Error())
	}

	// Start go rutine eth suscription to "NewVersion" event of all repos addresses
	fmt.Println(params.InfoLog + "Go rutine NewVersion event subscription...")
	go eth.SubscribeNewVersion(ethClient, dc, discordChannel, repos) 

	// Start go rutine eth suscription to "NewRepo" event of Registry 
 	fmt.Println(params.InfoLog + "Go rutine NewRepo event subscription...")
	go eth.SubscribeNewRepo(ethClient, dc, discordChannel) 

	fmt.Println(params.InfoLog + "Application successfully started")

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
