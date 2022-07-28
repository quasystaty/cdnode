package keeper

import (
	"context"
	"errors"

	"github.com/cdbo/cdnode/x/coinmaster/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	minters := k.Minters(ctx)
	if minters != types.DefaultMinters {
		if msg.Creator != minters {
			return nil, errors.New("unauthorized account")
		}
	}

	coins := sdk.NewCoins(msg.Amount)

	if !IsDenomWhiteListed(coins[0].Denom) {
		return nil, errors.New("unauthorized denom")
	}

	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx,
		addr,
		types.ModuleName,
		coins)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return nil, err
	}

	return &types.MsgBurnResponse{}, nil
}