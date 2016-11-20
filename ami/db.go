package ami

// DBDel Delete DB entry.
func DBDel(socket *Socket, actionID, family, key string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "DBDel",
		"ActionID": actionID,
		"Family":   family,
		"Key":      key,
	})
}

// DBDelTree delete DB tree.
func DBDelTree(socket *Socket, actionID, family, key string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "DBDelTree",
		"ActionID": actionID,
		"Family":   family,
		"Key":      key,
	})
}

// DBPut puts DB entry.
func DBPut(socket *Socket, actionID, family, key, val string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "DBPut",
		"ActionID": actionID,
		"Family":   family,
		"Key":      key,
		"Val":      val,
	})
}

// DBGet gets DB Entry.
func DBGet(socket *Socket, actionID, family, key string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "DBGet",
		"ActionID": actionID,
		"Family":   family,
		"Key":      key,
	})
}
