package ami

import "context"

// SIPNotify sends a SIP notify
func SIPNotify(ctx context.Context, client Client, actionID string, channel string, variable string) (Response, error) {
	return send(ctx, client, "SIPnotify", actionID, map[string]string{
		"Channel":  channel,
		"Variable": variable,
	})
}

// SIPPeers lists SIP peers in text format with details on current status.
// Peerlist will follow as separate events, followed by a final event called PeerlistComplete
func SIPPeers(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "SIPpeers", actionID, "PeerEntry", "PeerlistComplete")
}

// SIPPeerStatus show the status of one or all of the sip peers.
func SIPPeerStatus(ctx context.Context, client Client, actionID string, peer string) ([]Response, error) {
	if peer == "" {
		return requestList(ctx, client, "SIPpeerstatus", actionID, "PeerEntry", "PeerlistComplete")
	}
	return requestList(ctx, client, "SIPpeerstatus", actionID, "PeerEntry", "PeerlistComplete", map[string]string{
		"Peer": peer,
	})
}

// SIPQualifyPeer qualify SIP peers.
func SIPQualifyPeer(ctx context.Context, client Client, actionID string, peer string) (Response, error) {
	return send(ctx, client, "SIPqualifypeer", actionID, map[string]string{
		"Peer": peer,
	})
}

// SIPShowPeer shows one SIP peer with details on current status.
func SIPShowPeer(ctx context.Context, client Client, actionID string, peer string) (Response, error) {
	return send(ctx, client, "SIPshowpeer", actionID, map[string]string{
		"Peer": peer,
	})
}

// SIPShowRegistry shows SIP registrations (text format).
func SIPShowRegistry(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "SIPshowregistry", actionID, "RegistrationEntry", "RegistrationsComplete")
}
