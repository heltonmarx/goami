package ami

import "strconv"

// Agents lists agents and their status.
func Agents(socket *Socket, actionID string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "Agents",
		"ActionID": actionID,
	})
}

// AgentLogoff sets an agent as no longer logged in.
func AgentLogoff(socket *Socket, actionID, agent string, soft bool) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "AgentLogoff",
		"ActionID": actionID,
		"Agent":    agent,
		"Soft":     strconv.FormatBool(soft),
	})
}
