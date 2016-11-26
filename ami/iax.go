package ami

// IAXpeerlist show IAX channels network statistics.
func IAXpeerlist(client Client, actionID string) (Response, error) {
	return send(client, "IAXpeerlist", actionID, nil)
}

// IAXpeers list IAX peers.
func IAXpeers(client Client, actionID string) (Response, error) {
	return send(client, "IAXpeers", actionID, nil)
}

// IAXregistry show IAX registrations.
func IAXregistry(client Client, actionID string) (Response, error) {
	return send(client, "IAXregistry", actionID, nil)
}
