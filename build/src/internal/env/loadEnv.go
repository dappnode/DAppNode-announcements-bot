package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (gethRpc string, discordToken string, discordChannel string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

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

	return gethRpc, discordToken, discordChannel
}