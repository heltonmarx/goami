package ami

import "context"

// QueueAdd adds interface to queue.
func QueueAdd(ctx context.Context, client Client, actionID string, queueData QueueData) (Response, error) {
	return send(ctx, client, "QueueAdd", actionID, queueData)
}

// QueueLog adds custom entry in queue_log.
func QueueLog(ctx context.Context, client Client, actionID string, queueData QueueData) (Response, error) {
	return send(ctx, client, "QueueLog", actionID, queueData)
}

// QueuePause makes a queue member temporarily unavailable.
func QueuePause(ctx context.Context, client Client, actionID string, queueData QueueData) (Response, error) {
	return send(ctx, client, "QueuePause", actionID, queueData)
}

// QueuePenalty sets the penalty for a queue member.
func QueuePenalty(ctx context.Context, client Client, actionID string, queueData QueueData) (Response, error) {
	return send(ctx, client, "QueuePenalty", actionID, queueData)
}

// QueueReload reloads a queue, queues, or any sub-section of a queue or queues.
func QueueReload(ctx context.Context, client Client, actionID string, queueData QueueData) (Response, error) {
	return send(ctx, client, "QueueReload", actionID, queueData)
}

// QueueRemove removes interface from queue.
func QueueRemove(ctx context.Context, client Client, actionID string, queueData QueueData) (Response, error) {
	return send(ctx, client, "QueueRemove", actionID, queueData)
}

// QueueReset resets queue statistics.
func QueueReset(ctx context.Context, client Client, actionID, queue string) (Response, error) {
	return send(ctx, client, "QueueReset", actionID, QueueData{Queue: queue})
}

// QueueRule queues Rules.
func QueueRule(ctx context.Context, client Client, actionID, rule string) (Response, error) {
	return send(ctx, client, "QueueRule", actionID, map[string]string{
		"Rule": rule,
	})
}

// QueueStatus show queue status by member.
func QueueStatus(ctx context.Context, client Client, actionID, queue, member string) (Response, error) {
	return send(ctx, client, "QueueStatus", actionID, map[string]string{
		"Queue":  queue,
		"Member": member,
	})
}

// QueueStatuses show status all members in queue.
func QueueStatuses(ctx context.Context, client Client, actionID, queue string) ([]Response, error) {
	return requestMultiEvent(ctx, client, "QueueStatus", actionID, []string{"QueueMember", "QueueEntry"}, "QueueStatusComplete", map[string]string{
		"Queue": queue,
	})
}

// QueueSummary show queue summary.
func QueueSummary(ctx context.Context, client Client, actionID, queue string) ([]Response, error) {
	return requestList(ctx, client, "QueueSummary", actionID, "QueueSummary", "QueueSummaryComplete", map[string]string{
		"Queue": queue,
	})
}

// QueueMemberRingInUse set the ringinuse value for a queue member.
func QueueMemberRingInUse(ctx context.Context, client Client, actionID, iface, ringInUse, queue string) (Response, error) {
	return send(ctx, client, "QueueMemberRingInUse", actionID, map[string]string{
		"Interface": iface,
		"RingInUse": ringInUse,
		"Queue":     queue,
	})
}
