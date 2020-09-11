package ami

import "context"

// PRIDebugFileSet set the file used for PRI debug message output.
func PRIDebugFileSet(ctx context.Context, client Client, actionID, filename string) (Response, error) {
	return send(ctx, client, "PRIDebugFileSet", actionID, map[string]string{
		"File": filename,
	})
}

// PRIDebugFileUnset disables file output for PRI debug messages.
func PRIDebugFileUnset(ctx context.Context, client Client, actionID string) (Response, error) {
	return send(ctx, client, "PRIDebugFileUnset", actionID, nil)
}

// PRIDebugSet set PRI debug levels for a span.
func PRIDebugSet(ctx context.Context, client Client, actionID, span, level string) (Response, error) {
	return send(ctx, client, "PRIDebugSet", actionID, map[string]string{
		"Span":  span,
		"Level": level,
	})
}

// PRIShowSpans show status of PRI spans.
func PRIShowSpans(ctx context.Context, client Client, actionID, span string) ([]Response, error) {
	return requestList(ctx, client, "PRIShowSpans", actionID, "PRIShowSpans", "PRIShowSpansComplete", map[string]string{
		"Span": span,
	})
}
