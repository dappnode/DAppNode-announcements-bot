package env

import (
	"log"

	"github.com/joho/godotenv"
)

type Empty struct{}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
/* 		packagePath := reflect.TypeOf(Empty{}).PkgPath()
		_, packageName := path.Split(packagePath)
		err := godotenv.Load(".env." + packageName) */
}