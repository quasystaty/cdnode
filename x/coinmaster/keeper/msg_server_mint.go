package keeper

import (
	"context"

	"github.com/cdbo/cdnode/x/coinmaster/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(msg.Amount))
	if err != nil {
		return nil, err
	}

	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName,
		addr,
		sdk.NewCoins(msg.Amount))
	if err != nil {
		return nil, err
	}

	return &types.MsgMintResponse{}, nil
}
