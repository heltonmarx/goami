package ami

// Atxfer attended transfer.
func Atxfer(client Client, actionID, channel, exten, context string) (Response, error) {
	return send(client, "Atxfer", actionID, map[string]string{
		"Channel": channel,
		"Exten":   exten,
		"Context": context,
	})
}

// CancelAtxfer cancel an attended transfer.
func CancelAtxfer(client Client, actionID, channel string) (Response, error) {
	return send(client, "CancelAtxfer", actionID, map[string]string{
		"Channel": channel,
	})
}
