package handlers

import (
	"noteworthy/assets/env"
	"noteworthy/src/framework"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestMessageCreateHandlerShouldHandleMessage(t *testing.T) {
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
		//{
		//	name: "user posted message with correct prefix without args",
		//	message: &discordgo.MessageCreate{
		//		Message: &discordgo.Message{
		//			ChannelID: "exampleChannelID",
		//			Author: &discordgo.User{
		//				Username:      "exampleUsername",
		//				Discriminator: "1234",
		//				ID:            "userID",
		//			},
		//			Content: env.Prefix + "exampleCommand",
		//		},
		//	},
		//	expected: "Message created",
		//},
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
		//{
		//	name: "user posted message with correct prefix with args",
		//	message: &discordgo.MessageCreate{
		//		Message: &discordgo.Message{
		//			ChannelID: "exampleChannelID",
		//			Author: &discordgo.User{
		//				Username:      "exampleUsername",
		//				Discriminator: "1234",
		//				ID:            "userID",
		//			},
		//			Content: env.Prefix + "exampleCommand" + " exampleArgs",
		//		},
		//	},
		//	expected: "Message created",
		//},
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
		//{
		//	name: "user posted message with correct command",
		//	message: &discordgo.MessageCreate{
		//		Message: &discordgo.Message{
		//			ChannelID: "exampleChannelID",
		//			GuildID:   "exampleGuildID",
		//			Author: &discordgo.User{
		//				Username:      "exampleUsername",
		//				Discriminator: "1234",
		//				ID:            "exampleID",
		//			},
		//			Content: "!play",
		//		},
		//	},
		//	expected: "error getting guild",
		//},
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

func TestFillContext(t *testing.T) {
	//given
	userID := "exampleUserID"
	guildID := "exampleGuildID"
	channelID := "exampleChannelID"

	guild := &discordgo.Guild{
		ID: guildID,
		VoiceStates: []*discordgo.VoiceState{
			{
				GuildID: guildID,
			},
		},
	}
	channel := &discordgo.Channel{
		ID:      channelID,
		GuildID: guildID,
	}
	sess, _ := discordgo.New("")
	ctx := &framework.Context{}
	event := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			GuildID:   guildID,
			ChannelID: channelID,
			Author: &discordgo.User{
				ID: userID,
			},
		},
	}

	t.Run("should fail getting a guild", func(t *testing.T) {
		//when
		err := fillContext(ctx, sess, event)

		//then
		assert.EqualError(t, err, "error getting guild: state cache not found")
	})

	t.Run("should fail getting a text channel", func(t *testing.T) {
		//given
		err := sess.State.GuildAdd(guild)
		assert.NoError(t, err)

		//when
		err = fillContext(ctx, sess, event)

		//then
		assert.EqualError(t, err, "error getting text channel: state cache not found")
	})

	t.Run("should fail getting a voice state", func(t *testing.T) {
		//given
		err := sess.State.GuildAdd(guild)
		assert.NoError(t, err)
		err = sess.State.ChannelAdd(channel)
		assert.NoError(t, err)

		//when
		err = fillContext(ctx, sess, event)

		//then
		assert.EqualError(t, err, "You must be connected to the voice channel to use commands")
	})

	t.Run("should fail getting a voice channel", func(t *testing.T) {
		//given
		err := sess.State.GuildAdd(guild)
		assert.NoError(t, err)
		err = sess.State.ChannelAdd(channel)
		assert.NoError(t, err)
		guild.VoiceStates[0].UserID = userID

		//when
		err = fillContext(ctx, sess, event)

		//then
		assert.EqualError(t, err, "error getting voice channel: state cache not found")
	})

	t.Run("should succeed", func(t *testing.T) {
		//given
		err := sess.State.GuildAdd(guild)
		assert.NoError(t, err)
		err = sess.State.ChannelAdd(channel)
		assert.NoError(t, err)
		guild.VoiceStates[0].UserID = userID
		guild.VoiceStates[0].ChannelID = channelID

		//when
		err = fillContext(ctx, sess, event)

		//then
		assert.NoError(t, err)
	})
}
