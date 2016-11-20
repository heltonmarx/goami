package ami

// IAXpeerlist show IAX channels network statistics.
func IAXpeerlist(socket *Socket, actionID string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "IAXpeerlist",
		"ActionID": actionID,
	})
}

// IAXpeers list IAX peers.
func IAXpeers(socket *Socket, actionID string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "IAXpeers",
		"ActionID": actionID,
	})
}

// IAXregistry show IAX registrations.
func IAXregistry(socket *Socket, actionID string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "IAXregistry",
		"ActionID": actionID,
	})
}
