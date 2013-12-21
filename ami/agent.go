package ami

//  Agents
//      Lists agents and their status.
//
func Agents(socket *Socket, actionID string) ([]map[string]string, error) {
	return getMessageList(socket, "Agents", actionID, "AgentsEntry", "AgentsComplete")
}

//
//	AgentLogoff
//		Sets an agent as no longer logged in.
//
func AgentLogoff(socket *Socket, actionID, agent string, soft bool) (map[string]string, error) {
	//verify socket
	if !socket.Connected() {
		return nil, errors.New("Invalid socket")
	}

	// verify agent, soft and action ID
	if len(agent) == 0 || len(actionID) == 0 || len(soft) == 0 {
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
	err := sendCmd(socket, command)
	if err != nil {
		return nil, err
	}

	message, err := decode(socket)
	if (err != nil) || (message["ActionID"] != actionID) {
		return nil, err
	}
	return message, nil
}
