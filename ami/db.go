package ami

// DBDel Delete DB entry.
func DBDel(client Client, actionID, family, key string) (Response, error) {
	return send(client, "DBDel", actionID, dbData{
		Family: family,
		Key:    key,
	})
}

// DBDelTree delete DB tree.
func DBDelTree(client Client, actionID, family, key string) (Response, error) {
	return send(client, "DBDelTree", actionID, dbData{
		Family: family,
		Key:    key,
	})
}

// DBPut puts DB entry.
func DBPut(client Client, actionID, family, key, val string) (Response, error) {
	return send(client, "DBPut", actionID, dbData{
		Family: family,
		Key:    key,
		Value:  val,
	})
}

// DBGet gets DB Entry.
func DBGet(client Client, actionID, family, key string) (Response, error) {
	return send(client, "DBGet", actionID, dbData{
		Family: family,
		Key:    key,
	})
}

type dbData struct {
	Family string `ami:"Family"`
	Key    string `ami:"Key"`
	Value  string `ami:"Val,omitempty"`
}
