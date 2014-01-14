// Copyright 2014 Helton Marques
//
//	Use of this source code is governed by a LGPL
//	license that can be found in the LICENSE file.
//

package ami

import (
	"errors"
)

//
//	MailboxCount
//		Check Mailbox Message Count.
//
func MailboxCount(socket *Socket, actionID, mailbox string) (map[string]string, error) {
	if len(mailbox) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: MailboxCount",
		"\r\nActionID: ",
		actionID,
		"\r\nMailbox: ",
		mailbox,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}

//
//	MailboxStatus
//		Check Mailbox Message Count.
//
func MailboxStatus(socket *Socket, actionID, mailbox string) (map[string]string, error) {
	if len(mailbox) == 0 || len(actionID) == 0 {
		return nil, errors.New("Invalid parameters")
	}
	command := []string{
		"Action: MailboxStatus",
		"\r\nActionID: ",
		actionID,
		"\r\nMailbox: ",
		mailbox,
		"\r\n\r\n", // end of command
	}
	return getMessage(socket, command, actionID)
}
