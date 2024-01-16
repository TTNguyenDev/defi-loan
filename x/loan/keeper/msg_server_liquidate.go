package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"loan/x/loan/types"
)

func (k msgServer) Liquidate(goCtx context.Context, msg *types.MsgLiquidate) (*types.MsgLiquidateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgLiquidateResponse{}, nil
}
