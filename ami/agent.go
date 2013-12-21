package ami

//  Agents
//      Lists agents and their status.
//
func Agents(socket *Socket, actionID string) ([]map[string]string, error) {
	return getMessageList(socket, "Agents", actionID, "AgentsEntry", "AgentsComplete")
}
