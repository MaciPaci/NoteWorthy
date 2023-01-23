package framework

import (
	"github.com/bwmarrin/discordgo"
)

// Context hold important runtime variables
type Context struct {
	Session    *discordgo.Session
	User       *discordgo.User
	Message    *discordgo.Message
	BotMessage *Message
}

// NewContext creates a new Context
func NewContext(
	session *discordgo.Session,
	user *discordgo.User,
	message *discordgo.Message,
	botMessage *Message,
) *Context {
	return &Context{Session: session, User: user, Message: message, BotMessage: botMessage}
}
