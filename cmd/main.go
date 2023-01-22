package main

import (
	"NoteWorthy/assets/env"
	"NoteWorthy/utils/config"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func main() {
	conf, err := config.LoadConfig(env.ConfigFilePath)
	if err != nil {
		log.WithError(err).Fatal("Failed to load config")
	}
	_, err = discordgo.New("Bot" + conf.Token)
	if err != nil {
		log.WithError(err).Fatal("Failed to create bot")
	}
}
