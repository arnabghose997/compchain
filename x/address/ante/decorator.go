package ante

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type AddressFundReceiveRestriction struct{}

// AddressFundRecieveRestriction defines an Ante handler which restricts addresses ending with
// letter 's' from receiving any funds
func NewAddressFundReceiveRestriction() AddressFundReceiveRestriction {
	return AddressFundReceiveRestriction{}
}

func (afrr AddressFundReceiveRestriction) AnteHandle(
	ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler,
) (sdk.Context, error) {
	var totalMsgs []sdk.Msg

	msgs := tx.GetMsgs()

	// Gather any authz MsgExec sub-messages of type MsgSend and MsgMultiSend for
	// receiver address validation check
	for _, msg := range msgs {
		if authzExecMsg, ok := msg.(*authztypes.MsgExec); ok {
			subMessages, err := authzExecMsg.GetMessages()
			if err != nil {
				return sdk.Context{}, nil
			}
			totalMsgs = append(totalMsgs, subMessages...)
		} else {
			totalMsgs = append(totalMsgs, msg)
		}
	}

	if err := checkReceiverAddress(totalMsgs); err != nil {
		return sdk.Context{}, err
	}

	return next(ctx, tx, simulate)
}

// checkReceiverAddress loops through all transactions messages and look for any MsgSend or MsgMultiSend messages
// If found, check whether the reciever address ends with letter 's'
func checkReceiverAddress(msgs []sdk.Msg) error {
	for _, msg := range msgs {
		switch m := msg.(type) {
		case *banktypes.MsgSend:
			if strings.HasSuffix(m.ToAddress, "s") {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "%v is barred from receiving funds because it ends with `s`", m.ToAddress)
			}
		case *banktypes.MsgMultiSend:
			for _, reciever := range m.Outputs {
				if strings.HasSuffix(reciever.Address, "s") {
					return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "%v is barred from receiving funds because it ends with 's'", reciever.Address)
				}
			}
		}
	}
	return nil
}
