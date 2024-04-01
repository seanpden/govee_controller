package main

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/seanpden/govee_controller/pkg/cli_handler"
	clihandler "github.com/seanpden/govee_controller/pkg/cli_handler"
)

func handleEnvVar() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	APIKEY := os.Getenv("GOVEE_APIKEY")
	if APIKEY == "" {
		return "", errors.New("APIKEY not set")
	}
	return APIKEY, nil
}

func main() {
	APIKEY, err := handleEnvVar()
	if err != nil {
		log.Fatal(err)
	}
	clihandler.HandleCLI(APIKEY)
}
