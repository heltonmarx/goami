package ami

import "context"

// Bridge bridges two channels already in the PBX.
func Bridge(ctx context.Context, client Client, actionID, channel1, channel2 string, tone string) (Response, error) {
	return send(ctx, client, "Bridge", actionID, map[string]string{
		"Channel1": channel1,
		"Channel2": channel2,
		"Tone":     tone,
	})
}

// BlindTransfer blind transfer channel(s) to the given destination.
func BlindTransfer(ctx context.Context, client Client, actionID, channel, context, extension string) (Response, error) {
	return send(ctx, client, "BlindTransfer", actionID, map[string]string{
		"Channel": channel,
		"Context": context,
		"Exten":   extension,
	})
}

// BridgeDestroy destroy a bridge.
func BridgeDestroy(ctx context.Context, client Client, actionID, bridgeUniqueID string) (Response, error) {
	return send(ctx, client, "BridgeDestroy", actionID, map[string]string{
		"BridgeUniqueid": bridgeUniqueID,
	})
}

// BridgeInfo get information about a bridge.
func BridgeInfo(ctx context.Context, client Client, actionID, bridgeUniqueID string) (Response, error) {
	return send(ctx, client, "BridgeInfo", actionID, map[string]string{
		"BridgeUniqueid": bridgeUniqueID,
	})
}

// BridgeKick kick a channel from a bridge.
func BridgeKick(ctx context.Context, client Client, actionID, bridgeUniqueID, channel string) (Response, error) {
	params := map[string]string{
		"Channel": channel,
	}
	if bridgeUniqueID != "" {
		params["BridgeUniqueid"] = bridgeUniqueID
	}
	return send(ctx, client, "BridgeKick", actionID, params)
}

// BridgeList get a list of bridges in the system.
func BridgeList(ctx context.Context, client Client, actionID, bridgeType string) (Response, error) {
	return send(ctx, client, "BridgeList", actionID, map[string]string{
		"BridgeType": bridgeType,
	})
}

// BridgeTechnologyList list available bridging technologies and their statuses.
func BridgeTechnologyList(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "BridgeTechnologyList", actionID,
		"BridgeTechnologyListItem", "BridgeTechnologyListComplete")
}

// BridgeTechnologySuspend suspend a bridging technology.
func BridgeTechnologySuspend(ctx context.Context, client Client, actionID, bridgeTechnology string) (Response, error) {
	return send(ctx, client, "BridgeTechnologySuspend", actionID, map[string]string{
		"BridgeTechnology": bridgeTechnology,
	})
}

// BridgeTechnologyUnsuspend unsuspend a bridging technology.
func BridgeTechnologyUnsuspend(ctx context.Context, client Client, actionID, bridgeTechnology string) (Response, error) {
	return send(ctx, client, "BridgeTechnologyUnsuspend", actionID, map[string]string{
		"BridgeTechnology": bridgeTechnology,
	})
}
