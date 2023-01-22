package handlers

import (
	"github.com/bwmarrin/discordgo"
)

// Register adds event handlers to Discord session
func Register(s *discordgo.Session) {
	s.AddHandler(Ready)
	s.AddHandler(MessageCreate)
	s.AddHandler(MessageReactionAdded)
}
