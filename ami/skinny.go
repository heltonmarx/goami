package ami

import "context"

// SKINNYdevices lists SKINNY devices (text format).
// Lists Skinny devices in text format with details on current status.
// Devicelist will follow as separate events,
// followed by a final event called DevicelistComplete.
func SKINNYdevices(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "SKINNYdevices", actionID, "DeviceEntry", "DevicelistComplete")
}

// SKINNYlines lists SKINNY lines (text format).
// Lists Skinny lines in text format with details on current status.
// Linelist will follow as separate events,
// followed by a final event called LinelistComplete.
func SKINNYlines(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "SKINNYlines", actionID, "LineEntry", "LinelistComplete")
}

// SKINNYshowdevice show SKINNY device (text format).
// Show one SKINNY device with details on current status.
func SKINNYshowdevice(ctx context.Context, client Client, actionID, device string) (Response, error) {
	return send(ctx, client, "SKINNYshowdevice", actionID, map[string]string{
		"Device": device,
	})
}

// SKINNYshowline shows SKINNY line (text format).
// Show one SKINNY line with details on current status.
func SKINNYshowline(ctx context.Context, client Client, actionID, line string) (Response, error) {
	return send(ctx, client, "SKINNYshowline", actionID, map[string]string{
		"Line": line,
	})
}
