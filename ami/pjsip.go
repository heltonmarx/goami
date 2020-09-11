package ami

import "context"

// PJSIPNotify send NOTIFY to either an endpoint, an arbitrary URI, or inside a SIP dialog.
func PJSIPNotify(ctx context.Context, client Client, actionID, endpoint, uri, variable string) (Response, error) {
	params := map[string]string{
		"Variable": variable,
	}
	if endpoint != "" {
		params["Endpoint"] = endpoint
	}
	if uri != "" {
		params["URI"] = uri
	}
	return send(ctx, client, "PJSIPNotify", actionID, params)
}

// PJSIPQualify qualify a chan_pjsip endpoint.
func PJSIPQualify(ctx context.Context, client Client, actionID, endpoint string) (Response, error) {
	return send(ctx, client, "PJSIPQualify", actionID, map[string]string{
		"Endpoint": endpoint,
	})
}

// PJSIPRegister register an outbound registration.
func PJSIPRegister(ctx context.Context, client Client, actionID, registration string) (Response, error) {
	return send(ctx, client, "PJSIPRegister", actionID, map[string]string{
		"Registration": registration,
	})
}

// PJSIPUnregister unregister an outbound registration.
func PJSIPUnregister(ctx context.Context, client Client, actionID, registration string) (Response, error) {
	return send(ctx, client, "PJSIPUnregister", actionID, map[string]string{
		"Registration": registration,
	})
}

// PJSIPShowEndpoint detaill listing of an endpoint and its objects.
func PJSIPShowEndpoint(ctx context.Context, client Client, actionID, endpoint string) ([]Response, error) {
	return requestList(ctx, client, "PJSIPShowEndpoint", actionID, "EndpointDetail", "EndpointDetailComplete", map[string]string{
		"Endpoint": endpoint,
	})
}

// PJSIPShowEndpoints list pjsip endpoints.
func PJSIPShowEndpoints(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "PJSIPShowEndpoints", actionID, "EndpointList", "EndpointListComplete")
}

// PJSIPShowRegistrationInboundContactStatuses lists ContactStatuses for PJSIP inbound registrations.
func PJSIPShowRegistrationInboundContactStatuses(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "PJSIPShowRegistrationInboundContactStatuses", actionID, "ContactStatusDetail", "ContactStatusDetailComplete")
}

// PJSIPShowRegistrationsInbound lists PJSIP inbound registrations.
func PJSIPShowRegistrationsInbound(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "PJSIPShowRegistrationsInbound", actionID, "InboundRegistrationDetail", "InboundRegistrationDetailComplete")
}

// PJSIPShowRegistrationsOutbound lists PJSIP outbound registrations.
func PJSIPShowRegistrationsOutbound(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "PJSIPShowRegistrationsOutbound", actionID, "OutboundRegistrationDetail", "OutboundRegistrationDetailComplete")
}

// PJSIPShowResourceLists displays settings for configured resource lists.
func PJSIPShowResourceLists(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "PJSIPShowResourceLists", actionID, "ResourceListDetail", "ResourceListDetailComplete")
}

// PJSIPShowSubscriptionsInbound list of inbound subscriptions.
func PJSIPShowSubscriptionsInbound(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "PJSIPShowSubscriptionsInbound", actionID, "InboundSubscriptionDetail", "InboundSubscriptionDetailComplete")
}

// PJSIPShowSubscriptionsOutbound list of outbound subscriptions.
func PJSIPShowSubscriptionsOutbound(ctx context.Context, client Client, actionID string) ([]Response, error) {
	return requestList(ctx, client, "PJSIPShowSubscriptionsOutbound", actionID, "OutboundSubscriptionDetail", "OutboundSubscriptionDetailComplete")
}
