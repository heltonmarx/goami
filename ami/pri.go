package ami

// PRIDebugFileSet set the file used for PRI debug message output.
func PRIDebugFileSet(client Client, actionID, filename string) (Response, error) {
	return send(client, "PRIDebugFileSet", actionID, map[string]string{
		"File": filename,
	})
}

// PRIDebugFileUnset disables file output for PRI debug messages.
func PRIDebugFileUnset(client Client, actionID string) (Response, error) {
	return send(client, "PRIDebugFileUnset", actionID, nil)
}

// PRIDebugSet set PRI debug levels for a span.
func PRIDebugSet(client Client, actionID, span, level string) (Response, error) {
	return send(client, "PRIDebugSet", actionID, map[string]string{
		"Span":  span,
		"Level": level,
	})
}

// PRIShowSpans show status of PRI spans.
func PRIShowSpans(client Client, actionID, span string) ([]Response, error) {
	return requestList(client, "PRIShowSpans", actionID, "PRIShowSpans", "PRIShowSpansComplete", map[string]string{
		"Span": span,
	})
}
