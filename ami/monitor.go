package ami

import (
	"errors"
)

//	Monitor
//		Monitor a channel.
//		This action may be used to record the audio on a specified channel.
//
func Monitor(socket *Socket, actionID, channel, file, format string, mix bool) (map[string]string, error) {
	if len(actionID) == 0 || len(channel) == 0 ||
		len(file) == 0 || len(format) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	m := map[bool]string{false: "false", true: "true"}
	command := []string{
		"Action: Monitor",
		"\r\nActionID: ",
		actionID,
		"\r\nChannel: ",
		channel,
		"\r\nFile: ",
		file,
		"\r\nFormat: ",
		format,
		"\r\nMix: ",
		m[mix],
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}
