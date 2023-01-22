package handlers

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestMessageReactionAddedHandlerShouldRunSuccessfully(t *testing.T) {
	//given
	messageReactionAdd := &discordgo.MessageReactionAdd{
		MessageReaction: &discordgo.MessageReaction{
			UserID:    "exampleUserID",
			MessageID: "exampleMessageID",
			Emoji: discordgo.Emoji{
				Name: "exampleEmojiName",
			},
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
			", MessageID: "+messageReactionAdd.MessageID+
			", Emoji: "+messageReactionAdd.Emoji.Name)
}
