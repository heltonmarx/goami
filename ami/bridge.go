package ami

// Bridge bridges two channels already in the PBX.
func Bridge(client Client, actionID, channel1, channel2 string, tone string) (Response, error) {
	return send(client, "Bridge", actionID, map[string]string{
		"Channel1": channel1,
		"Channel2": channel2,
		"Tone":     tone,
	})
}

// BlindTransfer blind transfer channel(s) to the given destination.
func BlindTransfer(client Client, actionID, channel, context, extension string) (Response, error) {
	return send(client, "BlindTransfer", actionID, map[string]string{
		"Channel": channel,
		"Context": context,
		"Exten":   extension,
	})
}

// BridgeDestroy destroy a bridge.
func BridgeDestroy(client Client, actionID, bridgeUniqueID string) (Response, error) {
	return send(client, "BridgeDestroy", actionID, map[string]string{
		"BridgeUniqueid": bridgeUniqueID,
	})
}

// BridgeInfo get information about a bridge.
func BridgeInfo(client Client, actionID, bridgeUniqueID string) (Response, error) {
	return send(client, "BridgeInfo", actionID, map[string]string{
		"BridgeUniqueid": bridgeUniqueID,
	})
}

// BridgeKick kick a channel from a bridge.
func BridgeKick(client Client, actionID, bridgeUniqueID, channel string) (Response, error) {
	return send(client, "BridgeKick", actionID, map[string]string{
		"BridgeUniqueid": bridgeUniqueID,
		"Channel":        channel,
	})
}

// BridgeList get a list of bridges in the system.
func BridgeList(client Client, actionID, bridgeType string) (Response, error) {
	return send(client, "BridgeList", actionID, map[string]string{
		"BridgeType": bridgeType,
	})
}

// BridgeTechnologyList list available bridging technologies and their statuses.
func BridgeTechnologyList(client Client, actionID string) ([]Response, error) {
	return requestList(client, "BridgeTechnologyList", actionID,
		"BridgeTechnologyListItem", "BridgeTechnologyListComplete")
}

// BridgeTechnologySuspend suspend a bridging technology.
func BridgeTechnologySuspend(client Client, actionID, bridgeTechnology string) (Response, error) {
	return send(client, "BridgeTechnologySuspend", actionID, map[string]string{
		"BridgeTechnology": bridgeTechnology,
	})
}

// BridgeTechnologyUnsuspend unsuspend a bridging technology.
func BridgeTechnologyUnsuspend(client Client, actionID, bridgeTechnology string) (Response, error) {
	return send(client, "BridgeTechnologyUnsuspend", actionID, map[string]string{
		"BridgeTechnology": bridgeTechnology,
	})
}
