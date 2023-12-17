package ante

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
	msgs := tx.GetMsgs()

	// Loop through all transactions messages and look for any MsgSend or MsgMultiSend messages
	// If found, check whether the reciever address ends with letter 's'
	for _, msg := range msgs {
		switch m := msg.(type) {
		case *banktypes.MsgSend:
			if strings.HasSuffix(m.ToAddress, "s") {
				return sdk.Context{}, fmt.Errorf("%v is barred from receiving funds because it ends with `s`", m.ToAddress)
			}
		case *banktypes.MsgMultiSend:
			for _, reciever := range m.Outputs {
				if strings.HasSuffix(reciever.Address, "s") {
					return sdk.Context{}, fmt.Errorf("%v is barred from receiving funds because it ends with 's'", reciever.Address)
				}
			}
		}
	}

	return next(ctx, tx, simulate)
}
