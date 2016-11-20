package ami

// MailboxCount checks Mailbox Message Count.
func MailboxCount(socket *Socket, actionID, mailbox string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "MailboxCount",
		"ActionID": actionID,
		"Mailbox":  mailbox,
	})
}

// MailboxStatus checks Mailbox Message Count.
func MailboxStatus(socket *Socket, actionID, mailbox string) (map[string]string, error) {
	return sendCommand(socket, map[string]string{
		"Action":   "MailboxStatus",
		"ActionID": actionID,
		"Mailbox":  mailbox,
	})
}
