package ami

import "context"

const (
	iaxEvent    = "PeerEntry"
	iaxComplete = "PeerlistComplete"
)

// IAXnetstats show IAX channels network statistics.
func IAXnetstats(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "IAXnetstats", actionID, iaxEvent, iaxComplete)
}

// IAXpeerlist show IAX channels network statistics.
func IAXpeerlist(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "IAXpeerlist", actionID, iaxEvent, iaxComplete)
}

// IAXpeers list IAX peers.
func IAXpeers(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "IAXpeers", actionID, iaxEvent, iaxComplete)
}

// IAXregistry show IAX registrations.
func IAXregistry(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "IAXregistry", actionID, iaxEvent, iaxComplete)
}
