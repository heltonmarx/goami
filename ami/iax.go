package ami

const (
	iaxEvent    = "PeerEntry"
	iaxComplete = "PeerlistComplete"
)

// IAXnetstats show IAX channels network statistics.
func IAXnetstats(client Client, actionID string) ([]Response, error) {
	return requestList(client, "IAXnetstats", actionID, iaxEvent, iaxComplete)
}

// IAXpeerlist show IAX channels network statistics.
func IAXpeerlist(client Client, actionID string) ([]Response, error) {
	return requestList(client, "IAXpeerlist", actionID, iaxEvent, iaxComplete)
}

// IAXpeers list IAX peers.
func IAXpeers(client Client, actionID string) ([]Response, error) {
	return requestList(client, "IAXpeers", actionID, iaxEvent, iaxComplete)
}

// IAXregistry show IAX registrations.
func IAXregistry(client Client, actionID string) ([]Response, error) {
	return requestList(client, "IAXregistry", actionID, iaxEvent, iaxComplete)
}
