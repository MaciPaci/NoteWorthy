package handlers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErrUserNotInVoiceChannel(t *testing.T) {
	//given
	err := errors.New("some error")

	//when
	e := NewErrUserNotInVoiceChannel(err)

	//then
	assert.EqualError(t, e, "You must be connected to the voice channel to use commands")
}
