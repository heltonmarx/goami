package ami

// FAXStats responds with fax statistics.
func FAXStats(client Client, actionID string) (Response, error) {
	return send(client, "FAXStats", actionID, nil)
}

// FAXSession responds with a detailed description of a single FAX session.
func FAXSession(client Client, actionID string, sessionNumber string) (Response, error) {
	return send(client, "FAXSession", actionID, sessionNumber)
}

// FAXSessions lists active FAX sessions.
func FAXSessions(client Client, actionID string) ([]Response, error) {
	return requestList(client, "FAXSessions", actionID, "FAXSessionsEntry", "FAXSessionsComplete")
}
