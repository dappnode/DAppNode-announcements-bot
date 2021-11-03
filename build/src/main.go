package main

import (
	"announcements-bot/discord"
	"announcements-bot/eth"
	"fmt"
	"time"
)

/* func init() {
	fmt.Println("Setting up DAppNode announcements bot")

	environment := os.Getenv("token")
	if environment == "TEST" {
		viper.SetConfigFile(".env")
	} else if environment == "PRODUCTION" {

	} else {

	}

} */

var gethRpc string = "wss://mainnet.infura.io/ws/v3/e6c920580178424bbdf6dde266bfb5bd"
  
func main() {
	fmt.Println("Starting...")

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
	discord, err := discord.GetDiscordSession()
	if err != nil {
		err := fmt.Errorf("unable to open discord session: %w", err)
		fmt.Println(err)
	}
	
	// Start go rutine eth suscription to "NewVersion" event of all repos addresses
	fmt.Println("Go rutine NewVersion event subscription...")
	go eth.SubscribeNewVersion(ethClient, discord, repos)

	// Start go rutine eth suscription to "NewRepo" event of Registry 
	fmt.Println("Go rutine NewRepo event subscription...")
	go eth.SubscribeNewRepo(ethClient, discord)

	time.Sleep(10000000*1000000)
}




  

