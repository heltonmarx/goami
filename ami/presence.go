package ami

// PresenceState check presence state.
func PresenceState(client Client, actionID, provider string) (Response, error) {
	return send(client, "PresenceState", actionID, map[string]string{
		"Provider": provider,
	})
}

// PresenceStateList list the current known presence states.
func PresenceStateList(client Client, actionID string) ([]Response, error) {
	return requestList(client, "PresenceStateList", actionID,
		"Agents", "PresenceStateListComplete")
}
