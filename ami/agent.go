// Copyright 2014 Helton Marques
//
//	Use of this source code is governed by a LGPL
//	license that can be found in the LICENSE file.
//

package ami

import (
	"errors"
)

//  Agents
//      Lists agents and their status.
//
func Agents(socket *Socket, actionID string) ([]map[string]string, error) {
	command, err := getCommand("Agents", actionID)
	if err != nil {
		return nil, err
	}
	return getMessageList(socket, command, actionID,
		"AgentsEntry", "AgentsComplete")
}

//
//	AgentLogoff
//		Sets an agent as no longer logged in.
//
func AgentLogoff(socket *Socket, actionID, agent string, soft bool) (map[string]string, error) {
	if len(agent) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}

	s := map[bool]string{false: "false", true: "true"}
	command := []string{
		"Action: AgentLogoff",
		"\r\nActionID: ",
		actionID,
		"\r\nAgent: ",
		agent,
		"\r\nSoft: ",
		s[soft],
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}
