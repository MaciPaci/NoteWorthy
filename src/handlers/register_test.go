package handlers

import (
	"bytes"
	"os"
	"testing"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var session *discordgo.Session

func TestMain(m *testing.M) {
	session = &discordgo.Session{
		Token: "exampleToken",
		State: &discordgo.State{
			Ready: discordgo.Ready{
				User: &discordgo.User{
					ID: "botID",
				},
			},
		},
	}
	Register(session)
	exitVal := m.Run()
	os.Exit(exitVal)
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
