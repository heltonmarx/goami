package ami

// QueueAdd adds interface to queue.
func QueueAdd(client Client, actionID string, queueData QueueData) (Response, error) {
	return send(client, "QueueAdd", actionID, queueData)
}

// QueueLog adds custom entry in queue_log.
func QueueLog(client Client, actionID string, queueData QueueData) (Response, error) {
	return send(client, "QueueLog", actionID, queueData)
}

// QueuePause makes a queue member temporarily unavailable.
func QueuePause(client Client, actionID string, queueData QueueData) (Response, error) {
	return send(client, "QueuePause", actionID, queueData)
}

// QueuePenalty sets the penalty for a queue member.
func QueuePenalty(client Client, actionID string, queueData QueueData) (Response, error) {
	return send(client, "QueuePenalty", actionID, queueData)
}

// QueueReload reloads a queue, queues, or any sub-section of a queue or queues.
func QueueReload(client Client, actionID string, queueData QueueData) (Response, error) {
	return send(client, "QueueReload", actionID, queueData)
}

// QueueRemove removes interface from queue.
func QueueRemove(client Client, actionID string, queueData QueueData) (Response, error) {
	return send(client, "QueueRemove", actionID, queueData)
}

// QueueReset resets queue statistics.
func QueueReset(client Client, actionID, queue string) (Response, error) {
	return send(client, "QueueReset", actionID, QueueData{Queue: queue})
}

// QueueRule queues Rules.
func QueueRule(client Client, actionID, rule string) (Response, error) {
	return send(client, "QueueRule", actionID, map[string]string{
		"Rule": rule,
	})
}

// QueueStatus show queue status by member.
func QueueStatus(client Client, actionID, queue, member string) (Response, error) {
	return send(client, "QueueStatus", actionID, map[string]string{
		"Queue":  queue,
		"Member": member,
	})
}

<<<<<<< HEAD
// QueueStatuses show status all members in queue.
=======
// QueueStatuses show queue status all mebers.
>>>>>>> a0db663d9a4a4536d3707fbebd9f0aac63c2c081
func QueueStatuses(client Client, actionID, queue string) ([]Response, error) {
	return requestList(client, "QueueStatus", actionID, "QueueMember", "QueueStatusComplete", map[string]string{
		"Queue": queue,
	})
}

<<<<<<< HEAD

=======
>>>>>>> a0db663d9a4a4536d3707fbebd9f0aac63c2c081
// QueueSummary show queue summary.
func QueueSummary(client Client, actionID, queue string) ([]Response, error) {
	return requestList(client, "QueueSummary", actionID, "QueueSummary", "QueueSummaryComplete", map[string]string{
		"Queue": queue,
	})
}
