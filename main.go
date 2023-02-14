package main

import (
	"fmt"
	"github.com/dimaglushkov/contriseg/internal"
	"log"
)

const (
	usernameEnvVar = "GITHUB_USERNAME"
	tokenEnvVar    = "GITHUB_TOKEN"
)

func run() error {
	env, err := internal.GetEnvVariables(usernameEnvVar, tokenEnvVar)
	if err != nil {
		return fmt.Errorf("error getting env variables: %w", err)
	}

	cal, err := internal.GetContributions(env[usernameEnvVar], env[tokenEnvVar])
	if err != nil {
		return fmt.Errorf("error getting contributions stats from github: %w", err)
	}

	if err := internal.GenContributionsSegmentsGIF(cal, "contriseg.gif"); err != nil {
		return fmt.Errorf("error generating .gif: %w", err)
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
