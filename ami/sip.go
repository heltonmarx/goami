package ami

// SIPNotify sends a SIP notify
func SIPNotify(client Client, actionID string, channel string, variable string) (Response, error) {
	return send(client, "SIPnotify", actionID, map[string]string{
		"Channel":  channel,
		"Variable": variable,
	})
}

// SIPPeers lists SIP peers in text format with details on current status.
// Peerlist will follow as separate events, followed by a final event called PeerlistComplete
func SIPPeers(client Client, actionID string) ([]Response, error) {
	return requestList(client, "SIPpeers", actionID, "PeerEntry", "PeerlistComplete")
}

// SIPQualifyPeer qualify SIP peers.
func SIPQualifyPeer(client Client, actionID string, peer string) (Response, error) {
	return send(client, "SIPqualifypeer", actionID, map[string]string{
		"Peer": peer,
	})
}

// SIPShowPeer shows one SIP peer with details on current status.
func SIPShowPeer(client Client, actionID string, peer string) (Response, error) {
	return send(client, "SIPshowpeer", actionID, map[string]string{
		"Peer": peer,
	})
}

// SIPShowRegistry shows SIP registrations (text format).
func SIPShowRegistry(client Client, actionID string) ([]Response, error) {
	return requestList(client, "SIPshowregistry", actionID, "RegistrationEntry", "RegistrationsComplete")
}
