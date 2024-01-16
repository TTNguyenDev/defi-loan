package keeper

import (
	"context"
	"loan/x/loan/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RepayLoan(goCtx context.Context, msg *types.MsgRepayLoan) (*types.MsgRepayLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if loan.State != "approved" {
		return nil, errorsmod.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}
	lender, _ := sdk.AccAddressFromBech32(loan.Lender)
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	if msg.Creator != loan.Borrower {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot repay: not a borrower")
	}
	amount, _ := sdk.ParseCoinsNormalized(loan.Amount)
	fee, _ := sdk.ParseCoinsNormalized(loan.Fee)
	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral)
	// Why don't we just combine 2 numbers and then send at once????
	err := k.bankKeeper.SendCoins(ctx, borrower, lender, amount)
	if err != nil {
		return nil, err
	}
	err = k.bankKeeper.SendCoins(ctx, borrower, lender, fee)
	if err != nil {
		return nil, err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, collateral)
	if err != nil {
		return nil, err
	}
	loan.State = "repayed"
	k.SetLoan(ctx, loan)
	return &types.MsgRepayLoanResponse{}, nil
}
