package ami

import (
	"context"
	"strconv"
)

// Agents lists agents and their status.
func Agents(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "Agents", actionID, "Agents", "AgentsComplete")
}

// AgentLogoff sets an agent as no longer logged in.
func AgentLogoff(ctx context.Context, client Client, actionID, agent string, soft bool) (Response, error) {
	return send(ctx, client, "AgentLogoff", actionID, map[string]string{
		"Agent": agent,
		"Soft":  strconv.FormatBool(soft),
	})
}
