package ami

// MailboxCount checks Mailbox Message Count.
func MailboxCount(client Client, actionID, mailbox string) (Response, error) {
	return send(client, "MailboxCount", actionID, map[string]string{
		"Mailbox": mailbox,
	})
}

// MailboxStatus checks Mailbox Message Count.
func MailboxStatus(client Client, actionID, mailbox string) (Response, error) {
	return send(client, "MailboxStatus", actionID, map[string]string{
		"Mailbox": mailbox,
	})
}

// MWIDelete delete selected mailboxes.
func MWIDelete(client Client, actionID, mailbox string) (Response, error) {
	return send(client, "MWIDelete", actionID, map[string]string{
		"Mailbox": mailbox,
	})
}

// MWIGet get selected mailboxes with message counts.
func MWIGet(client Client, actionID, mailbox string) (Response, error) {
	return send(client, "MWIGet", actionID, map[string]string{
		"Mailbox": mailbox,
	})
}

// MWIUpdate update the mailbox message counts.
func MWIUpdate(client Client, actionID, mailbox, oldMessages, newMessages string) (Response, error) {
	return send(client, "MWIUpdate", actionID, map[string]string{
		"Mailbox":     mailbox,
		"OldMessages": oldMessages,
		"NewMessages": newMessages,
	})
}
