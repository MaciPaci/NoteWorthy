package commands

import (
	"noteworthy/src/framework"

	log "github.com/sirupsen/logrus"
)

// Play is a handler for play command
func Play(ctx framework.Context) {
	log.Info("Play command invoked")
	_, err := ctx.Session.ChannelVoiceJoin(ctx.Guild.ID, ctx.VoiceChanel.ID, false, true)
	if err != nil {
		return
	}
	//vc.Speaking(true)
	//vc.Speaking(false)
	//vc.Disconnect()
}
