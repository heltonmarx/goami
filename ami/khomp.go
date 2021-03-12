package ami

import "context"

// KSendSMS sends a SMS using KHOMP device.
func KSendSMS(ctx context.Context, client Client, actionID string, data KhompSMSData) (Response, error) {
	return send(ctx, client, "KSendSMS", actionID, data)
}
