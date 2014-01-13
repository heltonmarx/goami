package ami

//	IAXpeerlist
//		Show IAX channels network statistics.
//
func IAXpeerlist(socket *Socket, actionID string) ([]map[string]string, error) {
	command, err := getCommand("IAXpeerlist", actionID)
	if err != nil {
		return nil, err
	}
	return getMessageList(socket, command, actionID,
		"PeerEntry", "PeerlistComplete")
}

//	IAXpeers
//		List IAX peers.
//
func IAXpeers(socket *Socket, actionID string) ([]map[string]string, error) {
	command, err := getCommand("IAXpeers", actionID)
	if err != nil {
		return nil, err
	}
	return getMessageList(socket, command, actionID,
		"PeerEntry", "PeerlistComplete")
}

//	IAXregistry
//		Show IAX registrations.
//
func IAXregistry(socket *Socket, actionID string) ([]map[string]string, error) {
	command, err := getCommand("IAXregistry", actionID)
	if err != nil {
		return nil, err
	}
	return getMessageList(socket, command, actionID,
		"RegistryEntry", "RegistrationsComplete")
}
