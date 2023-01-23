package framework

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestNewContext(t *testing.T) {
	//given
	session := &discordgo.Session{}
	user := &discordgo.User{}
	message := &discordgo.Message{}
	botMessage := &Message{}
	expectedCtx := &Context{
		Session:    session,
		User:       user,
		Message:    message,
		BotMessage: botMessage,
	}

	//when
	ctx := NewContext(session, user, message, botMessage)

	//then
	assert.Equal(t, expectedCtx, ctx)
}
