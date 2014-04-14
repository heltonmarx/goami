// Copyright 2014 Helton Marques
//
//	Use of this source code is governed by a LGPL
//	license that can be found in the LICENSE file.
//

package ami

//	AOCMessage
//		Generate an Advice of Charge message on a channel.
//
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
	Channel                   string
	ChannelPrefix             string
	MsgType                   string
	ChargeType                string
	UnitAmount                string
	UnitType                  string
	CurrencyName              string
	CurrencyAmount            string
	CurrencyMultiplier        string
	TotalType                 string
	AocBillingId              string
	ChargingAssociationId     string
	ChargingAssociationNumber string
	ChargingAssociationPlan   string
}

//
//	Originate
//		Originate a call.
//		Generates an outgoing call to a Extension/Context/Priority or Application/Data
//
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
//		Async - Set to true for fast origination.
//		Codecs - Comma-separated list of codecs to use for this call.
//
type OriginateData struct {
	Channel     string
	Exten       string
	Context     string
	Priority    int
	Application string
	Data        string
	Timeout     int
	Callerid    string
	Variable    string
	Account     string
	Async       string
	Codecs      string
}

//	QueueData
//		used in functions:
//		QueueAdd, QueueLog, QueuePause,
//		QueuePenalty, QueueReload, QueueRemove,
//		QueueReset

type QueueData struct {
	Queue          string
	Interface      string
	Penalty        string
	Paused         string
	MemberName     string
	StateInterface string
	Event          string
	Uniqueid       string
	Message        string
	Reason         string
	Members        string
	Rules          string
	Parameters     string
}

//
//	KhompSMSData
//
type KhompSMSData struct {
	Device       string
	Destination  string
	Confirmation bool
	Message      string
}
