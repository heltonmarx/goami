package ami

// ConfbridgeList lists all users in a particular ConfBridge conference.
func ConfbridgeList(client Client, actionID string, conference string) ([]Response, error) {
	return requestList(client, "ConfbridgeList", actionID, "ConfbridgeList", "ConfbridgeListComplete", map[string]string{
		"Conference": conference,
	})
}

// ConfbridgeListRooms lists data about all active conferences.
func CConfbridgeKickonfbridgeListRooms(client Client, actionID string) ([]Response, error) {
	return requestList(client, "ConfbridgeListRooms", actionID, "ConfbridgeListRooms", "ConfbridgeListRoomsComplete")
}

// ConfbridgeMute mutes a specified user in a specified conference.
func ConfbridgeMute(client Client, actionID string, conference string, channel string) (Response, error) {
	return send(client, "ConfbridgeMute", actionID, map[string]string{
		"Conference": conference,
		"Channel":    channel,
	})
}

// ConfbridgeUnmute unmutes a specified user in a specified conferene.
func ConfbridgeUnmute(client Client, actionID string, conference string, channel string) (Response, error) {
	return send(client, "ConfbridgeUnmute", actionID, map[string]string{
		"Conference": conference,
		"Channel":    channel,
	})
}

// ConfbridgeKick removes a specified user from a specified conference.
func ConfbridgeKick(client Client, actionID string, conference string, channel string) (Response, error) {
	return send(client, "ConfbridgeKick", actionID, map[string]string{
		"Conference": conference,
		"Channel":    channel,
	})
}

// ConfbridgeLock locks a specified conference.
func ConfbridgeLock(client Client, actionID string, conference string, channel string) (Response, error) {
	return send(client, "ConfbridgeLock", actionID, map[string]string{
		"Conference": conference,
	})
}

// ConfbridgeUnlock unlocks a specified conference.
func ConfbridgeUnlock(client Client, actionID string, conference string, channel string) (Response, error) {
	return send(client, "ConfbridgeUnlock", actionID, map[string]string{
		"Conference": conference,
	})
}

// ConfbridgeSetSingleVideoSrc sets a conference user as the single video source distributed to all other video-capable participants.
func ConfbridgeSetSingleVideoSrc(client Client, actionID string, conference string, channel string) (Response, error) {
	return send(client, "ConfbridgeSetSingleVideoSrc", actionID, map[string]string{
		"Conference": conference,
		"Channel":    channel,
	})
}

// ConfbridgeStartRecord starts a recording in the context of given conference and creates a file with the name specified by recordFile
func ConfbridgeStartRecord(client Client, actionID string, conference string, recordFile string) (Response, error) {
	return send(client, "ConfbridgeStartRecord", actionID, map[string]string{
		"Conference": conference,
		"RecordFile": recordFile,
	})
}

// ConfbridgeStopRecord stops a recording pertaining to the given conference
func ConfbridgeStopRecord(client Client, actionID string, conference string) (Response, error) {
	return send(client, "ConfbridgeStopRecord", actionID, map[string]string{
		"Conference": conference,
	})
}
