package handlers

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// Register adds event handlers to Discord session
func Register(s *discordgo.Session) {
	s.AddHandler(Ready)
	s.AddHandler(MessageCreate)
	s.AddHandler(MessageReactionAdded)
}

// Ready is a handler for READY event
func Ready(session *discordgo.Session, event *discordgo.Ready) {
	log.Info("Session started")
	log.Infof("Logged in as %v#%v", event.User.Username, event.User.Discriminator)
}

// MessageCreate is a handler for message creation event
func MessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
	log.Info("Message created")
	log.Infof("Details: Channel: %v, Author: %v#%v, Content: %v",
		event.ChannelID, event.Author.Username, event.Author.Discriminator, event.Content)
}

// MessageReactionAdded is a handler for adding reaction to a message
func MessageReactionAdded(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
	log.Info("Message reaction added")
	log.Infof("Details: UserID: %v, ChannelID: %v, MessageID: %v, Emoji: %v",
		event.UserID, event.ChannelID, event.MessageID, event.Emoji.Name)
}
