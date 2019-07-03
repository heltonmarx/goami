package ami

// MessageSend send an out of call message to an endpoint.
func MessageSend(client Client, actionID string, message MessageData) (Response, error) {
	return send(client, "MessageSend", actionID, message)
}
