package ami

import "context"

// PresenceState check presence state.
func PresenceState(ctx context.Context, client Client, actionID, provider string) (Response, error) {
	return send(ctx, client, "PresenceState", actionID, map[string]string{
		"Provider": provider,
	})
}

// PresenceStateList list the current known presence states.
func PresenceStateList(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "PresenceStateList", actionID,
		"Agents", "PresenceStateListComplete")
}
