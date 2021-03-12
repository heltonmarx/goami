package ami

import "context"

// ConfbridgeList lists all users in a particular ConfBridge conference.
func ConfbridgeList(ctx context.Context, client Client, actionID string, conference string) ([]Response, error) {
	return requestList(ctx, client, "ConfbridgeList", actionID, "ConfbridgeList", "ConfbridgeListComplete", map[string]string{
		"Conference": conference,
	})
}

// ConfbridgeListRooms lists data about all active conferences.
func ConfbridgeListRooms(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "ConfbridgeListRooms", actionID, "ConfbridgeListRooms", "ConfbridgeListRoomsComplete")
}

// ConfbridgeMute mutes a specified user in a specified conference.
func ConfbridgeMute(ctx context.Context, client Client, actionID string, conference string, channel string) (Response, error) {
	return send(ctx, client, "ConfbridgeMute", actionID, map[string]string{
		"Conference": conference,
		"Channel":    channel,
	})
}

// ConfbridgeUnmute unmutes a specified user in a specified conferene.
func ConfbridgeUnmute(ctx context.Context, client Client, actionID string, conference string, channel string) (Response, error) {
	return send(ctx, client, "ConfbridgeUnmute", actionID, map[string]string{
		"Conference": conference,
		"Channel":    channel,
	})
}

// ConfbridgeKick removes a specified user from a specified conference.
func ConfbridgeKick(ctx context.Context, client Client, actionID string, conference string, channel string) (Response, error) {
	return send(ctx, client, "ConfbridgeKick", actionID, map[string]string{
		"Conference": conference,
		"Channel":    channel,
	})
}

// ConfbridgeLock locks a specified conference.
func ConfbridgeLock(ctx context.Context, client Client, actionID string, conference string, channel string) (Response, error) {
	return send(ctx, client, "ConfbridgeLock", actionID, map[string]string{
		"Conference": conference,
	})
}

// ConfbridgeUnlock unlocks a specified conference.
func ConfbridgeUnlock(ctx context.Context, client Client, actionID string, conference string, channel string) (Response, error) {
	return send(ctx, client, "ConfbridgeUnlock", actionID, map[string]string{
		"Conference": conference,
	})
}

// ConfbridgeSetSingleVideoSrc sets a conference user as the single video source distributed to all other video-capable participants.
func ConfbridgeSetSingleVideoSrc(ctx context.Context, client Client, actionID string, conference string, channel string) (Response, error) {
	return send(ctx, client, "ConfbridgeSetSingleVideoSrc", actionID, map[string]string{
		"Conference": conference,
		"Channel":    channel,
	})
}

// ConfbridgeStartRecord starts a recording in the context of given conference and creates a file with the name specified by recordFile
func ConfbridgeStartRecord(ctx context.Context, client Client, actionID string, conference string, recordFile string) (Response, error) {
	params := map[string]string{
		"Conference": conference,
	}
	if recordFile != "" {
		params["RecordFile"] = recordFile
	}
	return send(ctx, client, "ConfbridgeStartRecord", actionID, params)
}

// ConfbridgeStopRecord stops a recording pertaining to the given conference
func ConfbridgeStopRecord(ctx context.Context, client Client, actionID string, conference string) (Response, error) {
	return send(ctx, client, "ConfbridgeStopRecord", actionID, map[string]string{
		"Conference": conference,
	})
}
