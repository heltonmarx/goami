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
