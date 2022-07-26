package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type PermissionnedGovDecorator struct{}

const WHITELISTED_ADDR = "birth18xs5ya3xk2kp6skgsx08j6pwrqayzcyep8w5yt"

func NewPermissionnedGovDecorator() PermissionnedGovDecorator {
	return PermissionnedGovDecorator{}
}

func (ante PermissionnedGovDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	msgs := tx.GetMsgs()
	for _, m := range msgs {
		msg, ok := m.(*govtypes.MsgSubmitProposal)
		if !ok {
			continue
		}
		if msg.Proposer != WHITELISTED_ADDR {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "proposer address %s is not valid", msg.Proposer)
		}

		msgVote, ok := m.(*govtypes.MsgVote)
		if !ok {
			continue
		}
		if msgVote.Voter != WHITELISTED_ADDR {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "voter address %s is not valid", msgVote.Voter)
		}
	}

	return next(ctx, tx, simulate)
}
