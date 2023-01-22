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
