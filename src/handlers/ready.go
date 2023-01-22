package handlers

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// Ready is a handler for READY event
func Ready(session *discordgo.Session, event *discordgo.Ready) {
	log.Info("Session started")
	log.Infof("Logged in as %v#%v", event.User.Username, event.User.Discriminator)
}
