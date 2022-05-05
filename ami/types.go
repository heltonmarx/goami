package ami

// AOCData holds the information to generate an Advice of Charge message on a channel.
//		Channel - Channel name to generate the AOC message on.
//		ChannelPrefix -	Partial channel prefix. By using this option one can match the beginning part of a channel name without having to put the entire name in.
//						For example if a channel name is SIP/snom-00000001 and this value is set to SIP/snom, then that channel matches and the message will be sent.
//						Note however that only the first matched channel has the message sent on it.
//
//		MsgType - Defines what type of AOC message to create, AOC-D or AOC-E
//			D
//			E
//
//		ChargeType - Defines what kind of charge this message represents.
//			NA
//			FREE
//			Currency
//			Unit
//
//		UnitAmount(0) -	This represents the amount of units charged. The ETSI AOC standard specifies that this value along with the optional UnitType value are entries in a list.
//						To accommodate this these values take an index value starting at 0 which can be used to generate this list of unit entries.
//						For Example, If two unit entires were required this could be achieved by setting the paramter UnitAmount(0)=1234 and UnitAmount(1)=5678.
//						Note that UnitAmount at index 0 is required when ChargeType=Unit, all other entries in the list are optional.
//
//		UnitType(0) -	Defines the type of unit. ETSI AOC standard specifies this as an integer value between 1 and 16, but this value is left open to accept any positive integer.
//						Like the UnitAmount parameter, this value represents a list entry and has an index parameter that starts at 0.
//		CurrencyName - Specifies the currency's name. Note that this value is truncated after 10 characters.
//		CurrencyAmount - Specifies the charge unit amount as a positive integer. This value is required when ChargeType==Currency.
//
//		CurrencyMultiplier - Specifies the currency multiplier. This value is required when ChargeType==Currency.
//			OneThousandth
//			OneHundredth
//			OneTenth
//			One
//			Ten
//			Hundred
//			Thousand
//
//		TotalType - Defines what kind of AOC-D total is represented.
//			Total
//			SubTotal
//
//		AOCBillingId - Represents a billing ID associated with an AOC-D or AOC-E message. Note that only the first 3 items of the enum are valid AOC-D billing IDs
//			Normal
//			ReverseCharge
//			CreditCard
//			CallFwdUnconditional
//			CallFwdBusy
//			CallFwdNoReply
//			CallDeflection
//			CallTransfer
//
//		ChargingAssociationId - 	Charging association identifier. This is optional for AOC-E and can be set to any value between -32768 and 32767
//		ChargingAssociationNumber -	Represents the charging association party number. This value is optional for AOC-E.
//		ChargingAssociationPlan - 	Integer representing the charging plan associated with the ChargingAssociationNumber.
//									The value is bits 7 through 1 of the Q.931 octet containing the type-of-number and numbering-plan-identification fields.
//
type AOCData struct {
	Channel                   string `ami:"Channel"`
	ChannelPrefix             string `ami:"ChannelPrefix"`
	MsgType                   string `ami:"MsgType"`
	ChargeType                string `ami:"ChargeType"`
	UnitAmount                string `ami:"UnitAmount(0)"`
	UnitType                  string `ami:"UnitType(0)"`
	CurrencyName              string `ami:"CurrencyName"`
	CurrencyAmount            string `ami:"CurrencyAmount"`
	CurrencyMultiplier        string `ami:"CurrencyMultiplier"`
	TotalType                 string `ami:"TotalType"`
	AOCBillingID              string `ami:"AOCBillingId"`
	ChargingAssociationID     string `ami:"ChargingAssociationId"`
	ChargingAssociationNumber string `ami:"ChargingAssociationNumber"`
	ChargingAssociationPlan   string `ami:"ChargingAssociationPlan"`
}

// OriginateData holds information used to originate outgoing calls.
//		Channel - Channel name to call.
//		Exten - Extension to use (requires Context and Priority)
//		Context - Context to use (requires Exten and Priority)
//		Priority - Priority to use (requires Exten and Context)
//		Application - Application to execute.
//		Data - Data to use (requires Application).
//		Timeout - How long to wait for call to be answered (in ms.).
//		CallerID - Caller ID to be set on the outgoing channel.
//		Variable - Channel variable to set, multiple Variable: headers are allowed.
//		Account - Account code.
//		EarlyMedia - Set to true to force call bridge on early media.
//		Async - Set to true for fast origination.
//		Codecs - Comma-separated list of codecs to use for this call.
//		ChannelId - Channel UniqueId to be set on the channel.
//		OtherChannelId - Channel UniqueId to be set on the second local channel.
type OriginateData struct {
	Channel        string   `ami:"Channel,omitempty"`
	Exten          string   `ami:"Exten,omitempty"`
	Context        string   `ami:"Context,omitempty"`
	Priority       int      `ami:"Priority,omitempty"`
	Application    string   `ami:"Application,omitempty"`
	Data           string   `ami:"Data,omitempty"`
	Timeout        int      `ami:"Timeout,omitempty"`
	CallerID       string   `ami:"CallerID,omitempty"`
	Variable       []string `ami:"Variable,omitempty"`
	Account        string   `ami:"Account,omitempty"`
	EarlyMedia     string   `ami:"EarlyMedia,omitempty"`
	Async          string   `ami:"Async,omitempty"`
	Codecs         string   `ami:"Codecs,omitempty"`
	ChannelID      string   `ami:"ChannelId,omitempty"`
	OtherChannelID string   `ami:"OtherChannelId,omitempty"`
}

// QueueData holds to queue calls.
// used in functions:
//  QueueAdd, QueueLog, QueuePause,
//  QueuePenalty, QueueReload, QueueRemove,
//  QueueReset
type QueueData struct {
	Queue          string `ami:"Queue,omitempty"`
	Interface      string `ami:"Interface,omitempty"`
	Penalty        string `ami:"Penalty,omitempty"`
	Paused         string `ami:"Paused,omitempty"`
	MemberName     string `ami:"MemberName,omitempty"`
	StateInterface string `ami:"StateInterface,omitempty"`
	Event          string `ami:"Event,omitempty"`
	UniqueID       string `ami:"UniqueID,omitempty"`
	Message        string `ami:"Message,omitempty"`
	Reason         string `ami:"Reason,omitempty"`
	Members        string `ami:"Members,omitempty"`
	Rules          string `ami:"Rules,omitempty"`
	Parameters     string `ami:"Parameters,omitempty"`
}

// KhompSMSData holds the Khomp SMS information.
type KhompSMSData struct {
	Device       string `ami:"Device"`
	Destination  string `ami:"Destination"`
	Confirmation bool   `ami:"Confirmation"`
	Message      string `ami:"Message"`
}

// CallData holds the call data to transfer.
type CallData struct {
	Channel       string `ami:"Channel"`
	ExtraChannel  string `ami:"ExtraChannel,omitempty"`
	Exten         string `ami:"Exten"`
	ExtraExten    string `ami:"ExtraExten,omitempty"`
	Context       string `ami:"Context"`
	ExtraContext  string `ami:"ExtraContext,omitempty"`
	Priority      string `ami:"Priority"`
	ExtraPriority string `ami:"ExtraPriority,omitempty"`
}

// ExtensionData holds the extension data to dialplan.
type ExtensionData struct {
	Context         string `ami:"Context"`
	Extension       string `ami:"Extension"`
	Priority        string `ami:"Priority,omitempty"`
	Application     string `ami:"Application,omitempty"`
	ApplicationData string `ami:"ApplicationData,omitempty"`
	Replace         string `ami:"Replace,omitempty"`
}

// MessageData holds the message data to message send command.
type MessageData struct {
	To         string `ami:"To"`
	From       string `ami:"From"`
	Body       string `ami:"Body"`
	Base64Body string `ami:"Base64Body,omitempty"`
	Variable   string `ami:"Variable"`
}

// UpdateConfigAction holds the params for an action in UpdateConfig AMI command.
//
// example
// 	actions := make([]ami.UpdateConfigAction, 0)
//	actions = append(actions, ami.UpdateConfigAction{
// 		Action:   "EmptyCat",
// 		Category: "test01",
// 	})
// 	actions = append(actions, ami.UpdateConfigAction{
// 		Action:   "Append",
// 		Category: "test01",
// 		Var:      "type",
// 		Value:    "peer",
// 	})
type UpdateConfigAction struct {
	Action   string `ami:"Action"`
	Category string `ami:"Category"`
	Var      string `ami:"Var,omitempty"`
	Value    string `ami:"Value,omitempty"`
}
