package handlers

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// MessageReactionAdded is a handler for adding reaction to a input
func MessageReactionAdded(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
	log.Info("Message reaction added")
	log.Infof("Details: UserID: %v, MessageID: %v, Emoji: %v",
		event.UserID, event.MessageID, event.Emoji.Name)
}
