package types

import (
	"fmt"
)

// IBCRoute defines the IBC route for the tunnel module
var _ RouteI = &IBCHookRoute{}

// NewIBCHookRoute creates a new IBCHookRoute instance. It is used to create the IBC hook route data
func NewIBCHookRoute(channelID, destinationContractAddress string) *IBCHookRoute {
	return &IBCHookRoute{
		ChannelID:                  channelID,
		DestinationContractAddress: destinationContractAddress,
	}
}

// Route defines the IBC route for the tunnel module
func (r *IBCHookRoute) ValidateBasic() error {
	// Validate the ChannelID format
	if !IsChannelIDFormat(r.ChannelID) {
		return fmt.Errorf("channel identifier is not in the format: `channel-{N}`")
	}
	return nil
}

// NewIBCHookPacketReceipt creates a new IBCHookPacketReceipt instance. It is used to create the IBC hook packet receipt data
func NewIBCHookPacketReceipt(sequence uint64) *IBCHookPacketReceipt {
	return &IBCHookPacketReceipt{
		Sequence: sequence,
	}
}
