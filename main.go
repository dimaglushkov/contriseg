package main

import (
	"fmt"
	"github.com/dimaglushkov/contriseg/internal"
	"github.com/dimaglushkov/contriseg/internal/image"
	"log"
)

const (
	usernameEnvVar = "GITHUB_USERNAME"
	tokenEnvVar    = "GITHUB_TOKEN"
	locationEnvVar = "TARGET_LOCATION"
)

func run() error {
	env, err := internal.GetEnvVariables(usernameEnvVar, tokenEnvVar, locationEnvVar)
	if err != nil {
		return fmt.Errorf("error getting env variables: %w", err)
	}

	cal, err := internal.GetContributions(env[usernameEnvVar], env[tokenEnvVar])
	if err != nil {
		return fmt.Errorf("error getting contributions stats from github: %w", err)
	}

	frames, err := image.BFSFrames(cal)
	if err != nil {
		return fmt.Errorf("error getting frames: %w", err)
	}

	return image.GenerateGIF(frames, env[locationEnvVar])
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
