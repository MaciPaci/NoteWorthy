package handlers

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func Register(s *discordgo.Session) {
	s.AddHandler(Ready)
	s.AddHandler(MessageCreate)
	s.AddHandler(MessageReactionAdded)
}

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	log.Info("Session started")
	log.Infof("Logged in as %v#%v", r.User.Username, r.User.Discriminator)
}

func MessageCreate(s *discordgo.Session, r *discordgo.MessageCreate) {
	log.Info("Message created")
	log.Infof("Details: Channel: %v, Author: %v, Content: %v", r.ChannelID, r.Author, r.Content)
}

func MessageReactionAdded(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	log.Info("Message reaction added")
	log.Infof("Details: UserID: %v, MessageID: %v, Emoji: %v", r.UserID, r.MessageID, r.Emoji)
}
