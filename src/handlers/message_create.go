package handlers

import (
	"fmt"
	"noteworthy/src/commands"
	"noteworthy/src/framework"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// MessageCreate is a handler for input creation event
func MessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
	log.Info("Message created")
	log.Infof("Details: Author: %v#%v, Content: %v",
		event.Author.Username, event.Author.Discriminator, event.Content)

	if messagePostedByBot(session, event) {
		log.Info("Message posted by the bot itself. Skipping.")
		return
	}

	message, err := framework.ContentToMessage(event.Content)
	if err != nil {
		log.Infof("%v. Skipping", err.Error())
		return
	}

	command := commands.GetCommand(message.Command)
	if command.IsNil() {
		log.Warnf("unknown command: %v", message.Command)
		_, _ = session.ChannelMessageSend(event.ChannelID, fmt.Sprintf("Unknown command: %s", message.Command))
		return
	}

	ctx := framework.NewContext(session, event.Author, event.Message, message)
	err = fillContext(ctx, session, event)
	if err != nil {
		switch err.(type) {
		case ErrUserNotInVoiceChannel:
			_, _ = session.ChannelMessageSend(event.ChannelID, err.Error())
			log.Error("User not in voice channel")
			return
		default:
			log.Error(err)
			return
		}
	}

	command(*ctx)
}

func messagePostedByBot(session *discordgo.Session, event *discordgo.MessageCreate) bool {
	return event.Author.ID == session.State.User.ID
}

func fillContext(ctx *framework.Context, session *discordgo.Session, event *discordgo.MessageCreate) error {
	guild, err := session.State.Guild(event.GuildID)
	if err != nil {
		return fmt.Errorf("error getting guild: %v", err)
	}

	textChannel, err := session.State.Channel(event.ChannelID)
	if err != nil {
		return fmt.Errorf("error getting text channel: %v", err)
	}

	voiceState, err := session.State.VoiceState(event.GuildID, event.Author.ID)
	if err != nil {
		return NewErrUserNotInVoiceChannel(err)
	}

	voiceChannel, err := session.State.Channel(voiceState.ChannelID)
	if err != nil {
		return fmt.Errorf("error getting voice channel: %v", err)
	}

	ctx.Guild = guild
	ctx.TextChannel = textChannel
	ctx.VoiceChanel = voiceChannel

	return nil
}
