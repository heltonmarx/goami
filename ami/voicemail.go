package ami

import "context"

// VoicemailRefresh tell asterisk to poll mailboxes for a change.
func VoicemailRefresh(ctx context.Context, client Client, actionID, context, mailbox string) (Response, error) {
	return send(ctx, client, "VoicemailRefresh", actionID, map[string]string{
		"Context": context,
		"Mailbox": mailbox,
	})
}

// VoicemailUsersList list all voicemail user information.
func VoicemailUsersList(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "VoicemailUsersList", actionID, "VoicemailUserEntry", "VoicemailUserEntryComplete")
}
