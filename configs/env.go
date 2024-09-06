package configs

import "github.com/joho/godotenv"

func Env(envsFileNames ...string) {
	envsFileNames = append(envsFileNames, ".env")

	godotenv.Load(envsFileNames...)
}
