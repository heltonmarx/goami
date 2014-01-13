package ami

//	IAXpeerlist
//		Show IAX channels network statistics.
//
func IAXpeerlist(socket *Socket, actionID string) ([]map[string]string, error) {
	return getMessageList(socket, "IAXpeerlist", actionID, "PeerEntry", "PeerlistComplete")
}

//	IAXpeers
//		List IAX peers.
//
func IAXpeers(socket *Socket, actionID string) ([]map[string]string, error) {
	return getMessageList(socket, "IAXpeers", actionID, "PeerEntry", "PeerlistComplete")
}

//	IAXregistry
//		Show IAX registrations.
//
func IAXregistry(socket *Socket, actionID string) ([]map[string]string, error) {
	return getMessageList(socket, "IAXregistry", actionID, "RegistryEntry", "RegistrationsComplete")
}
