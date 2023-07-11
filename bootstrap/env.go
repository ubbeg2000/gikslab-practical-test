package bootstrap

import "github.com/joho/godotenv"

func Env() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
}
