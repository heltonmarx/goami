package ami

import "fmt"

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
	data := dbData{
		Family: family,
		Key:    key,
	}

	responses, err := requestList(client, "DBGet", actionID, "DBGetResponse", "DBGetComplete", data)

	switch {
	case err != nil:
		return nil, err
	case len(responses) == 0:
		return nil, fmt.Errorf("there is no db entries: family:%s key:%s", family, key)
	}

	return responses[0], nil
}

type dbData struct {
	Family string `ami:"Family"`
	Key    string `ami:"Key"`
	Value  string `ami:"Val,omitempty"`
}
