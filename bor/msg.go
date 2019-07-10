package bor

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
)

var cdc = codec.New()

// CheckpointRoute represents rount in app
const BorProposeSpanRoute = "proposeSpan"

//
// Propose Span Msg
//

var _ sdk.Msg = &MsgProposeSpan{}

type MsgProposeSpan struct {
	StartBlock uint64 `json:"startBlock"`

	// Timestamp only exits to allow submission of multiple transactions without bringing in nonce
	TimeStamp uint64 `json:"timestamp"`
}

// NewMsgProposeSpan creates new propose span message
func NewMsgProposeSpan(startBlock uint64, timestamp uint64) MsgProposeSpan {
	return MsgProposeSpan{
		StartBlock: startBlock,
		TimeStamp:  timestamp,
	}
}

// Type returns message type
func (msg MsgProposeSpan) Type() string {
	return "checkpoint"
}

// Route returns route for message
func (msg MsgProposeSpan) Route() string {
	return BorProposeSpanRoute
}

// GetSigners returns address of the signer
func (msg MsgProposeSpan) GetSigners() []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, 1)
	addrs[0] = sdk.AccAddress(msg.Proposer.Bytes())
	return addrs
}

// GetSignBytes returns sign bytes for proposeSpan message type
func (msg MsgProposeSpan) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// ValidateBasic validates the message and returns error
func (msg MsgProposeSpan) ValidateBasic() sdk.Error {
	if msg.TimeStamp == 0 || msg.TimeStamp > uint64(time.Now().Unix()) {
		return hmCommon.ErrInvalidMsg(hmCommon.DefaultCodespace, "Invalid timestamp %d", msg.TimeStamp)
	}
	return nil
}
