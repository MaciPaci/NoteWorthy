package handlers

import (
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
		//TODO handle wrong command
		log.Warnf("unknown command: %v", message.Command)
		return
	}
	ctx := framework.NewContext(session, event.Author, event.Message, message)
	command(*ctx)

	//_ = message.ToMessageEmbed()
	//_, err = session.ChannelMessageSendEmbed(event.ChannelID, embed)
	//if err != nil {
	//	log.Errorf("failed to post message: %v", err.Error())
	//}
}

func messagePostedByBot(session *discordgo.Session, event *discordgo.MessageCreate) bool {
	return event.Author.ID == session.State.User.ID
}
