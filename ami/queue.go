// Copyright 2014 Helton Marques
//
//	Use of this source code is governed by a LGPL
//	license that can be found in the LICENSE file.
//

package ami

import (
	"errors"
)

var (
	errInvalidQueueParameters = errors.New("queue: Invalid parameters")
)

//	QueueAdd
//		Add interface to queue.
//
func QueueAdd(socket *Socket, actionID string, queueData QueueData) (map[string]string, error) {
	if len(actionID) == 0 {
		return nil, errInvalidQueueParameters
	}

	// verify struct parameters
	if len(queueData.Queue) == 0 || len(queueData.Interface) == 0 ||
		len(queueData.Penalty) == 0 || len(queueData.Paused) == 0 ||
		len(queueData.MemberName) == 0 || len(queueData.StateInterface) == 0 {
		return nil, errInvalidQueueParameters
	}

	command := []string{
		"Action: QueueAdd",
		"\r\nActionID: ",
		actionID,
		"\r\nQueue: ",
		queueData.Queue,
		"\r\nInterface: ",
		queueData.Interface,
		"\r\nPenalty: ",
		queueData.Penalty,
		"\r\nPaused: ",
		queueData.Paused,
		"\r\nMemberName: ",
		queueData.MemberName,
		"\r\nStateInterface: ",
		queueData.StateInterface,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	QueueLog
//		Adds custom entry in queue_log.
//
func QueueLog(socket *Socket, actionID string, queueData QueueData) (map[string]string, error) {
	if len(actionID) == 0 {
		return nil, errInvalidQueueParameters
	}

	// verify struct parameters
	if len(queueData.Queue) == 0 || len(queueData.Event) == 0 ||
		len(queueData.Event) == 0 || len(queueData.Uniqueid) == 0 ||
		len(queueData.Interface) == 0 || len(queueData.Message) == 0 {
		return nil, errInvalidQueueParameters
	}

	command := []string{
		"Action: QueueLog",
		"\r\nActionID: ",
		actionID,
		"\r\nQueue: ",
		queueData.Queue,
		"\r\nEvent: ",
		queueData.Event,
		"\r\nUniqueid: ",
		queueData.Uniqueid,
		"\r\nInterface: ",
		queueData.Interface,
		"\r\nMessage: ",
		queueData.Message,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	QueuePause
//		Makes a queue member temporarily unavailable.
//
func QueuePause(socket *Socket, actionID string, queueData QueueData) (map[string]string, error) {
	if len(actionID) == 0 {
		return nil, errInvalidQueueParameters
	}

	// verify struct parameters
	if len(queueData.Queue) == 0 || len(queueData.Interface) == 0 ||
		len(queueData.Paused) == 0 || len(queueData.Reason) == 0 {
		return nil, errInvalidQueueParameters
	}

	command := []string{
		"Action: QueuePause",
		"\r\nActionID: ",
		actionID,
		"\r\nQueue: ",
		queueData.Queue,
		"\r\nInterface: ",
		queueData.Interface,
		"\r\nPaused: ",
		queueData.Paused,
		"\r\nReason: ",
		queueData.Reason,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	QueuePenalty
//		Set the penalty for a queue member.
//
func QueuePenalty(socket *Socket, actionID string, queueData QueueData) (map[string]string, error) {
	if len(actionID) == 0 {
		return nil, errInvalidQueueParameters
	}

	// verify struct parameters
	if len(queueData.Queue) == 0 || len(queueData.Interface) == 0 || len(queueData.Penalty) == 0 {
		return nil, errInvalidQueueParameters
	}

	command := []string{
		"Action: QueuePenalty",
		"\r\nActionID: ",
		actionID,
		"\r\nQueue: ",
		queueData.Queue,
		"\r\nInterface: ",
		queueData.Interface,
		"\r\nPenalty: ",
		queueData.Penalty,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	QueueReload
//		Reload a queue, queues, or any sub-section of a queue or queues.
//
func QueueReload(socket *Socket, actionID string, queueData QueueData) (map[string]string, error) {
	if len(actionID) == 0 {
		return nil, errInvalidQueueParameters
	}

	// verify struct parameters
	if len(queueData.Queue) == 0 || len(queueData.Members) == 0 ||
		len(queueData.Rules) == 0 || len(queueData.Parameters) == 0 {
		return nil, errInvalidQueueParameters
	}

	command := []string{
		"Action: QueueReload",
		"\r\nActionID: ",
		actionID,
		"\r\nQueue: ",
		queueData.Queue,
		"\r\nMembers: ",
		queueData.Members,
		"\r\nRules: ",
		queueData.Rules,
		"\r\nParameters: ",
		queueData.Parameters,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	QueueRemove
//		Remove interface from queue.
//
func QueueRemove(socket *Socket, actionID string, queueData QueueData) (map[string]string, error) {
	if len(actionID) == 0 {
		return nil, errInvalidQueueParameters
	}

	// verify struct parameters
	if len(queueData.Queue) == 0 || len(queueData.Interface) == 0 {
		return nil, errInvalidQueueParameters
	}

	command := []string{
		"Action: QueueRemove",
		"\r\nActionID: ",
		actionID,
		"\r\nQueue: ",
		queueData.Queue,
		"\r\nInterface: ",
		queueData.Interface,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	QueueReset
//		Reset queue statistics.
//
func QueueReset(socket *Socket, actionID, queue string) (map[string]string, error) {
	if len(actionID) == 0 || len(queue) == 0 {
		return nil, errInvalidQueueParameters
	}
	command := []string{
		"Action: QueueReset",
		"\r\nActionID: ",
		actionID,
		"\r\nQueue: ",
		queue,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	QueueRule
//		Queue Rules.
//
func QueueRule(socket *Socket, actionID, rule string) (map[string]string, error) {
	if len(actionID) == 0 || len(rule) == 0 {
		return nil, errInvalidQueueParameters
	}
	command := []string{
		"Action: QueueRule",
		"\r\nActionID: ",
		actionID,
		"\r\nRule: ",
		rule,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	QueueStatus
//		Show queue status.
//
func QueueStatus(socket *Socket, actionID, queue, member string) (map[string]string, error) {
	if len(actionID) == 0 || len(queue) == 0 || len(member) == 0 {
		return nil, errInvalidQueueParameters
	}
	command := []string{
		"Action: QueueStatus",
		"\r\nActionID: ",
		actionID,
		"\r\nQueue: ",
		queue,
		"\r\nMember: ",
		member,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//	QueueSummary
//		Show queue summary.
//
func QueueSummary(socket *Socket, actionID, queue string) (map[string]string, error) {
	if len(actionID) == 0 || len(queue) == 0 {
		return nil, errInvalidQueueParameters
	}
	command := []string{
		"Action: QueueSummary",
		"\r\nActionID: ",
		actionID,
		"\r\nQueue: ",
		queue,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}
