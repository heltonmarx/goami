package ami

// FAXSession responds with a detailed description of a single FAX session.
func FAXSession(client Client, actionID, sessionNumber string) (Response, error) {
	return send(client, "FAXSession", actionID, map[string]string{
		"SessionNumber": sessionNumber,
	})
}

// FAXSessions list active FAX sessions.
func FAXSessions(client Client, actionID string) ([]Response, error) {
	return requestList(client, "FAXSessions", actionID, "FAXSessionsEntry", "FAXSessionsComplete")
}

// FAXStats responds with fax statistics.
func FAXStats(client Client, actionID string) (Response, error) {
	return send(client, "FAXStats", actionID, nil)
}
