package handlers

import (
	"bytes"
	"os"
	"testing"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var session *discordgo.Session

func TestMain(m *testing.M) {
	session = &discordgo.Session{
		Token: "exampleToken",
	}
	Register(session)
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestReadyHandlerShouldRunSuccessfully(t *testing.T) {
	//given
	ready := &discordgo.Ready{
		User: &discordgo.User{
			Username:      "exampleUsername",
			Discriminator: "1234",
		},
	}

	//when
	output := captureOutput(func() {
		Ready(session, ready)
	})

	//then
	assert.Contains(t, output, "Session started")
	assert.Contains(t, output, "Logged in as "+ready.User.Username+"#"+ready.User.Discriminator)
}

func TestMessageCreateHandlerShouldRunSuccessfully(t *testing.T) {
	//given
	messageCreate := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ChannelID: "exampleChannelID",
			Author: &discordgo.User{
				Username:      "exampleUsername",
				Discriminator: "1234",
			},
			Content: "exampleContent",
		},
	}

	//when
	output := captureOutput(func() {
		MessageCreate(session, messageCreate)
	})

	//then
	assert.Contains(t, output, "Message created")
	assert.Contains(t, output,
		"Details: "+
			"Channel: "+messageCreate.ChannelID+
			", Author: "+messageCreate.Author.Username+"#"+messageCreate.Author.Discriminator+
			", Content: "+messageCreate.Content)
}

func TestMessageReactionAddedHandlerShouldRunSuccessfully(t *testing.T) {
	//given
	messageReactionAdd := &discordgo.MessageReactionAdd{
		MessageReaction: &discordgo.MessageReaction{
			UserID:    "exampleUserID",
			MessageID: "exampleMessageID",
			Emoji: discordgo.Emoji{
				Name: "exampleEmojiName",
			},
			ChannelID: "exampleChannelID",
		},
	}

	//when
	output := captureOutput(func() {
		MessageReactionAdded(session, messageReactionAdd)
	})

	//then
	assert.Contains(t, output, "Message reaction added")
	assert.Contains(t, output,
		"Details: "+
			"UserID: "+messageReactionAdd.UserID+
			", ChannelID: "+messageReactionAdd.ChannelID+
			", MessageID: "+messageReactionAdd.MessageID+
			", Emoji: "+messageReactionAdd.Emoji.Name)
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
