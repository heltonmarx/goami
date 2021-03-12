package ami

import "context"

// FAXSession responds with a detailed description of a single FAX session.
func FAXSession(ctx context.Context, client Client, actionID, sessionNumber string) (Response, error) {
	return send(ctx, client, "FAXSession", actionID, map[string]string{
		"SessionNumber": sessionNumber,
	})
}

// FAXSessions list active FAX sessions.
func FAXSessions(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "FAXSessions", actionID, "FAXSessionsEntry", "FAXSessionsComplete")
}

// FAXStats responds with fax statistics.
func FAXStats(ctx context.Context, client Client, actionID string) (Response, error) {
	return send(ctx, client, "FAXStats", actionID, nil)
}
