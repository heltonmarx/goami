package ami

// MeetmeList lists all users in a particular MeetMe conference.
// Will follow as separate events, followed by a final event called MeetmeListComplete.
func MeetmeList(client Client, actionID, conference string) ([]Response, error) {
	return requestList(client, "MeetmeList", actionID, "MeetmeEntry", "MeetmeListComplete")
}

// MeetmeMute mute a Meetme user.
func MeetmeMute(client Client, actionID, meetme, usernum string) (Response, error) {
	return send(client, "MeetmeMute", actionID, map[string]string{
		"Meetme":  meetme,
		"Usernum": usernum,
	})
}

// MeetmeUnMute unmute a Meetme user.
func MeetmeUnMute(client Client, actionID, meetme, usernum string) (Response, error) {
	return send(client, "MeetmeUnMute", actionID, map[string]string{
		"Meetme":  meetme,
		"Usernum": usernum,
	})
}
