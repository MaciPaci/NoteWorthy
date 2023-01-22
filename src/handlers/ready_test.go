package handlers

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

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
