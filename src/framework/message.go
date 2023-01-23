package framework

import (
	"errors"
	"noteworthy/assets/env"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Message represents bot message
type Message struct {
	Prefix  string
	Command string
	Args    string
	Color   int
}

// Colors is a map of embed message colors
var Colors = map[string]int{
	"red":   0xb20000,
	"green": 0x00b200,
	"blue":  0x0000b2,
}

// ToMessageEmbed converts Message to MessageEmbed
func (m *Message) ToMessageEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       m.Command,
		Description: m.Args,
		Color:       m.Color,
	}
}

// ContentToMessage converts event content to Message
func ContentToMessage(content string) (*Message, error) {
	messagePrefix, messageContent := extractPrefix(content)

	if messagePrefix != env.Prefix {
		return &Message{}, errors.New("no prefix found")
	}

	messageCommand, messageArgs, err := extractCommandAndArgs(messageContent)
	if err != nil {
		return &Message{}, err
	}

	return &Message{
		Prefix:  messagePrefix,
		Command: messageCommand,
		Args:    messageArgs,
	}, nil
}

func extractPrefix(message string) (string, string) {
	messagePrefix := message[:len(env.Prefix)]
	messageContent := message[len(env.Prefix):]
	return messagePrefix, messageContent
}

func extractCommandAndArgs(message string) (string, string, error) {
	var messageCommand, messageArgs string
	if len(message) < 1 {
		return "", "", errors.New("no message content found")
	}
	messageContent := strings.SplitN(message, " ", 2)
	if len(messageContent) > 1 {
		messageArgs = messageContent[1]
	}
	messageCommand = messageContent[0]
	return messageCommand, messageArgs, nil
}
