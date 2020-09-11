package ami

import "context"

// MailboxCount checks Mailbox Message Count.
func MailboxCount(ctx context.Context, client Client, actionID, mailbox string) (Response, error) {
	return send(ctx, client, "MailboxCount", actionID, map[string]string{
		"Mailbox": mailbox,
	})
}

// MailboxStatus checks Mailbox Message Count.
func MailboxStatus(ctx context.Context, client Client, actionID, mailbox string) (Response, error) {
	return send(ctx, client, "MailboxStatus", actionID, map[string]string{
		"Mailbox": mailbox,
	})
}

// MWIDelete delete selected mailboxes.
func MWIDelete(ctx context.Context, client Client, actionID, mailbox string) (Response, error) {
	return send(ctx, client, "MWIDelete", actionID, map[string]string{
		"Mailbox": mailbox,
	})
}

// MWIGet get selected mailboxes with message counts.
func MWIGet(ctx context.Context, client Client, actionID, mailbox string) (Response, error) {
	return send(ctx, client, "MWIGet", actionID, map[string]string{
		"Mailbox": mailbox,
	})
}

// MWIUpdate update the mailbox message counts.
func MWIUpdate(ctx context.Context, client Client, actionID, mailbox, oldMessages, newMessages string) (Response, error) {
	return send(ctx, client, "MWIUpdate", actionID, map[string]string{
		"Mailbox":     mailbox,
		"OldMessages": oldMessages,
		"NewMessages": newMessages,
	})
}
