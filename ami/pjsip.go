package ami

// PJSIPNotify send NOTIFY to either an endpoint, an arbitrary URI, or inside a SIP dialog.
func PJSIPNotify(client Client, actionID, endpoint, uri, channel, variable string) (Response, error) {
	return send(client, "PJSIPNotify", actionID, map[string]string{
		"Endpoint": endpoint,
		"URI":      uri,
		"channel":  channel,
		"Variable": variable,
	})
}

// PJSIPQualify qualify a chan_pjsip endpoint.
func PJSIPQualify(client Client, actionID, endpoint string) (Response, error) {
	return send(client, "PJSIPQualify", actionID, map[string]string{
		"Endpoint": endpoint,
	})
}

// PJSIPRegister register an outbound registration.
func PJSIPRegister(client Client, actionID, registration string) (Response, error) {
	return send(client, "PJSIPRegister", actionID, map[string]string{
		"Registration": registration,
	})
}

// PJSIPUnregister unregister an outbound registration.
func PJSIPUnregister(client Client, actionID, registration string) (Response, error) {
	return send(client, "PJSIPUnregister", actionID, map[string]string{
		"Registration": registration,
	})
}

// PJSIPShowEndpoint detaill listing of an endpoint and its objects.
func PJSIPShowEndpoint(client Client, actionID, endpoint string) ([]Response, error) {
	return requestList(client, "PJSIPShowEndpoint", actionID, "EndpointDetail", "EndpointDetailComplete", map[string]string{
		"Endpoint": endpoint,
	})
}

// PJSIPShowEndpoints list pjsip endpoints.
func PJSIPShowEndpoints(client Client, actionID string) ([]Response, error) {
	return requestList(client, "PJSIPShowEndpoints", actionID, "EndpointList", "EndpointListComplete")
}

// PJSIPShowRegistrationInboundContactStatuses lists ContactStatuses for PJSIP inbound registrations.
func PJSIPShowRegistrationInboundContactStatuses(client Client, actionID string) ([]Response, error) {
	return requestList(client, "PJSIPShowRegistrationInboundContactStatuses", actionID, "ContactStatusDetail", "ContactStatusDetailComplete")
}

// PJSIPShowRegistrationsInbound lists PJSIP inbound registrations.
func PJSIPShowRegistrationsInbound(client Client, actionID string) ([]Response, error) {
	return requestList(client, "PJSIPShowRegistrationsInbound", actionID, "InboundRegistrationDetail", "InboundRegistrationDetailComplete")
}

// PJSIPShowRegistrationsOutbound lists PJSIP outbound registrations.
func PJSIPShowRegistrationsOutbound(client Client, actionID string) ([]Response, error) {
	return requestList(client, "PJSIPShowRegistrationsOutbound", actionID, "OutboundRegistrationDetail", "OutboundRegistrationDetailComplete")
}

// PJSIPShowResourceLists displays settings for configured resource lists.
func PJSIPShowResourceLists(client Client, actionID string) ([]Response, error) {
	return requestList(client, "PJSIPShowResourceLists", actionID, "ResourceListDetail", "ResourceListDetailComplete")
}

// PJSIPShowSubscriptionsInbound list of inbound subscriptions.
func PJSIPShowSubscriptionsInbound(client Client, actionID string) ([]Response, error) {
	return requestList(client, "PJSIPShowSubscriptionsInbound", actionID, "InboundSubscriptionDetail", "InboundSubscriptionDetailComplete")
}

// PJSIPShowSubscriptionsOutbound list of outbound subscriptions.
func PJSIPShowSubscriptionsOutbound(client Client, actionID string) ([]Response, error) {
	return requestList(client, "PJSIPShowSubscriptionsOutbound", actionID, "OutboundSubscriptionDetail", "OutboundSubscriptionDetailComplete")
}
