package main

import (
	"noteworthy/assets/env"
	"noteworthy/src/handlers"
	"noteworthy/utils/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func main() {
	conf, err := config.LoadConfig(env.ConfigFilePath)
	if err != nil {
		log.WithError(err).Fatal("failed to load config")
	}

	session, err := discordgo.New("Bot " + conf.Token)
	if err != nil {
		log.WithError(err).Fatal("failed to create the session")
	}

	handlers.Register(session)

	err = session.Open()
	if err != nil {
		log.WithError(err).Fatal("failed to open the session")
	}
	defer session.Close()

	log.Info("NoteWorthy is now running. Press CTRL-C to exit.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}
