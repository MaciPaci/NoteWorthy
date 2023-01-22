package handlers

import (
	"noteworthy/assets/env"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestMessageCreateHandlerShouldRunSuccessfully(t *testing.T) {
	testCases := []struct {
		name     string
		message  *discordgo.MessageCreate
		expected string
	}{
		{
			name: "user posted message without any prefix",
			message: &discordgo.MessageCreate{
				Message: &discordgo.Message{
					ChannelID: "exampleChannelID",
					Author: &discordgo.User{
						Username:      "exampleUsername",
						Discriminator: "1234",
						ID:            "userID",
					},
					Content: "exampleCommand",
				},
			},
			expected: "no prefix found",
		},
		{
			name: "user posted message with wrong prefix",
			message: &discordgo.MessageCreate{
				Message: &discordgo.Message{
					ChannelID: "exampleChannelID",
					Author: &discordgo.User{
						Username:      "exampleUsername",
						Discriminator: "1234",
						ID:            "userID",
					},
					Content: "@exampleCommand",
				},
			},
			expected: "no prefix found",
		},
		{
			name: "user posted message with correct prefix without args",
			message: &discordgo.MessageCreate{
				Message: &discordgo.Message{
					ChannelID: "exampleChannelID",
					Author: &discordgo.User{
						Username:      "exampleUsername",
						Discriminator: "1234",
						ID:            "userID",
					},
					Content: env.Prefix + "exampleCommand",
				},
			},
			expected: "Message created",
		},
		{
			name: "user posted message only with prefix",
			message: &discordgo.MessageCreate{
				Message: &discordgo.Message{
					ChannelID: "exampleChannelID",
					Author: &discordgo.User{
						Username:      "exampleUsername",
						Discriminator: "1234",
						ID:            "userID",
					},
					Content: env.Prefix,
				},
			},
			expected: "no message content found",
		},
		{
			name: "user posted message with correct prefix with args",
			message: &discordgo.MessageCreate{
				Message: &discordgo.Message{
					ChannelID: "exampleChannelID",
					Author: &discordgo.User{
						Username:      "exampleUsername",
						Discriminator: "1234",
						ID:            "userID",
					},
					Content: env.Prefix + "exampleCommand" + " exampleArgs",
				},
			},
			expected: "Message created",
		},
		{
			name: "message was posted by a bot",
			message: &discordgo.MessageCreate{
				Message: &discordgo.Message{
					ChannelID: "exampleChannelID",
					Author: &discordgo.User{
						Username:      "exampleUsername",
						Discriminator: "1234",
						ID:            "botID",
					},
					Content: "exampleContent",
				},
			},
			expected: "Message posted by the bot itself. Skipping.",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//when
			output := captureOutput(func() {
				MessageCreate(session, testCase.message)
			})

			//then
			assert.Contains(t, output, testCase.expected)
			assert.Contains(t, output,
				"Details: "+
					"Author: "+testCase.message.Author.Username+"#"+testCase.message.Author.Discriminator+
					", Content: "+testCase.message.Content)
		})
	}
}
