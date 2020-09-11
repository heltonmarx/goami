package ami

import "context"

// MeetmeList lists all users in a particular MeetMe conference.
// Will follow as separate events, followed by a final event called MeetmeListComplete.
func MeetmeList(ctx context.Context, client Client, actionID, conference string) ([]Response, error) {
	return requestList(ctx, client, "MeetmeList", actionID, "MeetmeEntry", "MeetmeListComplete")
}

// MeetmeMute mute a Meetme user.
func MeetmeMute(ctx context.Context, client Client, actionID, meetme, usernum string) (Response, error) {
	return send(ctx, client, "MeetmeMute", actionID, map[string]string{
		"Meetme":  meetme,
		"Usernum": usernum,
	})
}

// MeetmeUnMute unmute a Meetme user.
func MeetmeUnMute(ctx context.Context, client Client, actionID, meetme, usernum string) (Response, error) {
	return send(ctx, client, "MeetmeUnMute", actionID, map[string]string{
		"Meetme":  meetme,
		"Usernum": usernum,
	})
}

// MeetmeListRooms list active conferences.
func MeetmeListRooms(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "MeetmeListRooms", actionID, "MeetmeEntry", "MeetmeListRoomsComplete")
}
