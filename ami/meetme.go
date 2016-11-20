package ami

// MeetmeList lists all users in a particular MeetMe conference.
// Will follow as separate events, followed by a final event called MeetmeListComplete.
func MeetmeList(socket *Socket, actionID, conference string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":     "MeetmeList",
		"ActionID":   actionID,
		"Conference": conference,
	})
}

// MeetmeMute mute a Meetme user.
func MeetmeMute(socket *Socket, actionID, meetme, usernum string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "MeetmeMute",
		"ActionID": actionID,
		"Meetme":   meetme,
		"Usernum":  usernum,
	})
}

// MeetmeUnMute unmute a Meetme user.
func MeetmeUnMute(socket *Socket, actionID, meetme, usernum string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "MeetmeUnMute",
		"ActionID": actionID,
		"Meetme":   meetme,
		"Usernum":  usernum,
	})
}
