package env

import (
	"announcements-bot/params"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Empty struct{}

func LoadEnv() {
	goEnv := os.Getenv("GO_ENV")
	if (goEnv == "") {
		panic("GO_ENV is not set")
	}
	
	if(goEnv == "production") {
		/* err := godotenv.Load(".env")
		if err != nil {
			err := fmt.Errorf("unable to load env: %w", err)
			panic(params.ErrorLog + err.Error())
		} */
		return
	} else {
		err := godotenv.Load("test.env")
		if err != nil {
			err := fmt.Errorf("unable to load env: %w", err)
			panic(params.ErrorLog + err.Error())
		}
	}
}