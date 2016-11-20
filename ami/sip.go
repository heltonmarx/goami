package ami

// SIPNotify sends a SIP notify
func SIPNotify(socket *Socket, actionID string, channel string, variable string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "SIPnotify",
		"ActionID": actionID,
		"Channel":  channel,
		"Variable": variable,
	})
}

// SIPPeers lists SIP peers in text format with details on current status.
// Peerlist will follow as separate events, followed by a final event called PeerlistComplete
func SIPPeers(socket *Socket, actionID string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "SIPpeers",
		"ActionID": actionID,
	})
}

// SIPQualifyPeer qualify SIP peers.
func SIPQualifyPeer(socket *Socket, actionID string, peer string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "SIPqualifypeer",
		"ActionID": actionID,
		"Peer":     peer,
	})
}

// SIPShowPeer shows one SIP peer with details on current status.
func SIPShowPeer(socket *Socket, actionID string, peer string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "SIPshowpeer",
		"ActionID": actionID,
		"Peer":     peer,
	})
}

// SIPShowRegistry shows SIP registrations (text format).
func SIPShowRegistry(socket *Socket, actionID string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "SIPshowregistry",
		"ActionID": actionID,
	})
}
