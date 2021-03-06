// Code generated by 'banks apply'; DO NOT EDIT.

package resources

import (
	"github.com/autopilothq/banks/contract"
	"github.com/autopilothq/banks/protocol/marshalable"
	"github.com/autopilothq/banks/types"

	"github.com/hackathon/journeys/protocols"
	rodata "github.com/hackathon/journeys/resourcefile"

	// register resources and load their protocols
	_ "github.com/hackathon/journeys/resources/Journeys"
	protocolJourneys "github.com/hackathon/journeys/resources/Journeys/protocol"
)

func init() {
	lock, err := rodata.GetROLockfile()
	if err != nil {
		panic(err)
	}

	contract.Provides(
		lock.Resources,
		lock.Resource,
		protocols.Read,
		MsgIDFromReqBody,
	)
}

// MsgIDFromReqBody looks up a message ID for a request body type. It is used
// by banks dispatch.
func MsgIDFromReqBody(reqBody marshalable.Body) (types.MessageID, bool) {
	switch reqBody.(type) {

	case *protocolJourneys.GetTreeRequest:
		return types.MessageID(65538), true

	default:
		return types.MessageID(0), false
	}
}
