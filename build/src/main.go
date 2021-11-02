package main

import (
	"announcements-bot/discord"
	"announcements-bot/eth"
	"fmt"
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
  
func main() {
	// Connect to ETH node
	ethClient, err := eth.GetEthClient()
	if err != nil {
		fmt.Println(err)
	}

	go eth.NewRepoSub(ethClient)

	eth.NewVersionSub(ethClient, "0x448BfB454718f20941FE8a1bfA63a0024F21Ba50")

	// Start discord server
	discord.StartServer() 

	/* repos, err := file.GetRepos()
	if err != nil {
		fmt.Println(err)
	}  */
}




  

