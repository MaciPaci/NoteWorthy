package handlers

// ErrUserNotInVoiceChannel is an error returned when user calls command while not being in voice channel
type ErrUserNotInVoiceChannel struct {
	err error
}

// Error returns string error for ErrUserNotInVoiceChannel
func (e ErrUserNotInVoiceChannel) Error() string {
	return "You must be connected to the voice channel to use commands"
}

// NewErrUserNotInVoiceChannel created new ErrUserNotInVoiceChannel
func NewErrUserNotInVoiceChannel(err error) ErrUserNotInVoiceChannel {
	return ErrUserNotInVoiceChannel{err}
}
