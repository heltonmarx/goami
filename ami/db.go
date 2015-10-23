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
	errInvalidDBParameters = errors.New("DB: Invalid parameters")
)

//
//	DBDel
//		Delete DB entry.
//
func DBDel(socket *Socket, actionID, family, key string) (map[string]string, error) {
	if len(actionID) == 0 || len(family) == 0 || len(key) == 0 {
		return nil, errInvalidDBParameters
	}
	command := []string{
		"Action: DBDel",
		"\r\nActionID: ",
		actionID,
		"\r\nFamily: ",
		family,
		"\r\nKey: ",
		key,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	DBDelTree
//		Delete DB tree.
//
func DBDelTree(socket *Socket, actionID, family, key string) (map[string]string, error) {
	if len(actionID) == 0 || len(family) == 0 || len(key) == 0 {
		return nil, errInvalidDBParameters
	}
	command := []string{
		"Action: DBDelTree",
		"\r\nActionID: ",
		actionID,
		"\r\nFamily: ",
		family,
		"\r\nKey: ",
		key,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	DBPut
//		Put DB entry.
//
func DBPut(socket *Socket, actionID, family, key, val string) (map[string]string, error) {
	if len(actionID) == 0 || len(family) == 0 || len(key) == 0 {
		return nil, errInvalidDBParameters
	}
	command := []string{
		"Action: DBPut",
		"\r\nActionID: ",
		actionID,
		"\r\nFamily: ",
		family,
		"\r\nKey: ",
		key,
		"\r\nVal: ",
		val,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	DBGet
//		Get DB Entry.
//
func DBGet(socket *Socket, actionID, family, key string) (map[string]string, error) {
	if len(actionID) == 0 || len(family) == 0 || len(key) == 0 {
		return nil, errInvalidDBParameters
	}
	command := []string{
		"Action: DBGet",
		"\r\nActionID: ",
		actionID,
		"\r\nFamily: ",
		family,
		"\r\nKey: ",
		key,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}
