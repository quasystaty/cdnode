package keeper

import (
	permgovtypes "github.com/cdbo/cdnode/x/permgov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type PermissionnedGovDecorator struct {
	keeper Keeper
}

func NewPermissionnedGovDecorator(kpr Keeper) PermissionnedGovDecorator {
	return PermissionnedGovDecorator{
		keeper: kpr,
	}
}

func (ante PermissionnedGovDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	governor := ante.keeper.Governor(ctx)
	// Bypass the AnteHandler if the governor is not set
	if governor == permgovtypes.DefaultGovernor {
		return next(ctx, tx, simulate)
	}

	msgs := tx.GetMsgs()

	// Loops through all the messages and checks if the governor is allowed to execute the message
	for _, m := range msgs {
		msg, ok := m.(*govtypes.MsgSubmitProposal)
		if !ok {
			continue
		}
		if msg.Proposer != governor {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "proposer address %s is not valid", msg.Proposer)
		}

		msgVote, ok := m.(*govtypes.MsgVote)
		if !ok {
			continue
		}
		if msgVote.Voter != governor {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "voter address %s is not valid", msgVote.Voter)
		}
	}

	return next(ctx, tx, simulate)
}
