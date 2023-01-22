package input

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestMessageToMessageEmbed(t *testing.T) {
	//given
	message := Message{
		Prefix:  "examplePrefix",
		Command: "exampleCommand",
		Args:    "exampleArgs",
		Color:   0x00000,
	}
	expectedEmbed := &discordgo.MessageEmbed{
		Title:       message.Command,
		Description: message.Args,
		Color:       message.Color,
	}

	//when
	embed := message.ToMessageEmbed()

	//then
	assert.Equal(t, expectedEmbed, embed)
}

func TestContentToMessageShouldSuccessfullyCreateMessage(t *testing.T) {
	//given
	testCases := []struct {
		name            string
		content         string
		expectedMessage *Message
	}{
		{
			name:    "user posted a message with prefix and command",
			content: "!exampleCommand",
			expectedMessage: &Message{
				Prefix:  "!",
				Command: "exampleCommand",
			},
		},
		{
			name:    "user posted a message with prefix and command and args",
			content: "!exampleCommand exampleArgs",
			expectedMessage: &Message{
				Prefix:  "!",
				Command: "exampleCommand",
				Args:    "exampleArgs",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//when
			message, err := ContentToMessage(testCase.content)

			//then
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedMessage, message)
		})
	}
}

func TestContentToMessageShouldFailInCreatingMessage(t *testing.T) {
	//given
	testCases := []struct {
		name                 string
		content              string
		expectedErrorMessage string
		expectedMessage      *Message
	}{
		{
			name:                 "user posted message without any prefix",
			content:              "exampleContent",
			expectedErrorMessage: "no prefix found",
			expectedMessage:      &Message{},
		},
		{
			name:                 "user posted message with wrong prefix",
			content:              "@exampleContent",
			expectedErrorMessage: "no prefix found",
			expectedMessage:      &Message{},
		},
		{
			name:                 "user posted message only with prefix",
			content:              "!",
			expectedErrorMessage: "no message content found",
			expectedMessage:      &Message{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//when
			message, err := ContentToMessage(testCase.content)

			//then
			assert.Equal(t, err.Error(), testCase.expectedErrorMessage)
			assert.Equal(t, testCase.expectedMessage, message)
		})
	}
}
