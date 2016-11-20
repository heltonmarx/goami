package ami

// QueueAdd adds interface to queue.
func QueueAdd(socket *Socket, actionID string, queueData QueueData) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":         "QueueAdd",
		"ActionID":       actionID,
		"Queue":          queueData.Queue,
		"Interface":      queueData.Interface,
		"Penalty":        queueData.Penalty,
		"Paused":         queueData.Paused,
		"MemberName":     queueData.MemberName,
		"StateInterface": queueData.StateInterface,
	})
}

// QueueLog adds custom entry in queue_log.
func QueueLog(socket *Socket, actionID string, queueData QueueData) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":    "QueueLog",
		"ActionID":  actionID,
		"Queue":     queueData.Queue,
		"Event":     queueData.Event,
		"Uniqueid":  queueData.Uniqueid,
		"Interface": queueData.Interface,
		"Message":   queueData.Message,
	})
}

// QueuePause makes a queue member temporarily unavailable.
func QueuePause(socket *Socket, actionID string, queueData QueueData) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":    "QueuePause",
		"ActionID":  actionID,
		"Queue":     queueData.Queue,
		"Interface": queueData.Interface,
		"Paused":    queueData.Paused,
		"Reason":    queueData.Reason,
	})
}

// QueuePenalty sets the penalty for a queue member.
func QueuePenalty(socket *Socket, actionID string, queueData QueueData) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":    "QueuePenalty",
		"ActionID":  actionID,
		"Queue":     queueData.Queue,
		"Interface": queueData.Interface,
		"Penalty":   queueData.Penalty,
	})
}

// QueueReload reloads a queue, queues, or any sub-section of a queue or queues.
func QueueReload(socket *Socket, actionID string, queueData QueueData) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":     "QueueReload",
		"ActionID":   actionID,
		"Queue":      queueData.Queue,
		"Members":    queueData.Members,
		"Rules":      queueData.Rules,
		"Parameters": queueData.Parameters,
	})
}

// QueueRemove removes interface from queue.
func QueueRemove(socket *Socket, actionID string, queueData QueueData) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":    "QueueRemove",
		"ActionID":  actionID,
		"Queue":     queueData.Queue,
		"Interface": queueData.Interface,
	})
}

// QueueReset resets queue statistics.
func QueueReset(socket *Socket, actionID, queue string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "QueueReset",
		"ActionID": actionID,
		"Queue":    queue,
	})
}

// QueueRule queues Rules.
func QueueRule(socket *Socket, actionID, rule string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "QueueRule",
		"ActionID": actionID,
		"Rule":     rule,
	})
}

// QueueStatus show queue status.
func QueueStatus(socket *Socket, actionID, queue, member string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "QueueStatus",
		"ActionID": actionID,
		"Queue":    queue,
		"Member":   member,
	})
}

// QueueSummary show queue summary.
func QueueSummary(socket *Socket, actionID, queue string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "QueueSummary",
		"ActionID": actionID,
		"Queue":    queue,
	})
}
