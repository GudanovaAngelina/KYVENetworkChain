package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"github.com/KYVENetwork/chain/util"
	"github.com/KYVENetwork/chain/x/funders/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsTypes "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) defundLowestFunding(
	ctx sdk.Context,
	lowestFunding *types.Funding,
	fundingState *types.FundingState,
	poolId uint64,
) error {
	err := util.TransferFromModuleToAddress(k.bankKeeper, ctx, types.ModuleName, lowestFunding.FunderAddress, lowestFunding.Amount)
	if err != nil {
		return err
	}

	lowestFunding.SubtractAmount(lowestFunding.Amount)
	fundingState.SetInactive(lowestFunding)
	k.SetFunding(ctx, lowestFunding)

	// Emit a defund event.
	_ = ctx.EventManager().EmitTypedEvent(&types.EventDefundPool{
		PoolId:  poolId,
		Address: lowestFunding.FunderAddress,
		Amount:  lowestFunding.Amount,
	})
	return nil
}

// FundPool handles the logic to fund a pool.
// A funder is added to the active funders list with the specified amount
// If the funders list is full, it checks if the funder wants to fund
// more than the current lowest funder. If so, the current lowest funder
// will get their tokens back and removed form the active funders list.
func (k msgServer) FundPool(goCtx context.Context, msg *types.MsgFundPool) (*types.MsgFundPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Funder has to exist
	if !k.DoesFunderExist(ctx, msg.Creator) {
		return nil, errors.Wrapf(errorsTypes.ErrNotFound, types.ErrFunderDoesNotExist.Error(), msg.Creator)
	}

	// Pool has to exist
	err := k.poolKeeper.AssertPoolExists(ctx, msg.PoolId)
	if err != nil {
		return nil, err
	}

	// Get funding state for pool
	fundingState, found := k.GetFundingState(ctx, msg.PoolId)
	if !found {
		return nil, errors.Wrapf(errorsTypes.ErrNotFound, types.ErrFundingStateDoesNotExist.Error(), msg.PoolId)
	}

	// Check if funding already exists
	funding, found := k.GetFunding(ctx, msg.Creator, msg.PoolId)
	if found {
		// If so, update funding amount
		funding.AddAmount(msg.Amount)
		// If the amount per bundle is set, update it
		if msg.AmountPerBundle > 0 {
			funding.AmountPerBundle = msg.AmountPerBundle
		}
	} else {
		// If not, create new funding
		funding = types.Funding{
			FunderAddress:   msg.Creator,
			PoolId:          msg.PoolId,
			Amount:          msg.Amount,
			AmountPerBundle: msg.AmountPerBundle,
			TotalFunded:     0,
		}
	}
	fundingState.AddAmount(msg.Amount)

	params := k.GetParams(ctx)
	if funding.AmountPerBundle < params.MinFundingAmountPerBundle {
		return nil, errors.Wrapf(errorsTypes.ErrInvalidRequest, types.ErrAmountPerBundleTooLow.Error(), params.MinFundingAmountPerBundle)
	}
	if funding.Amount < params.MinFundingAmount {
		return nil, errors.Wrapf(errorsTypes.ErrInvalidRequest, types.ErrMinFundingAmount.Error(), params.MinFundingAmount)
	}

	var defunding *types.Funding = nil

	activeFundings := k.GetActiveFundings(ctx, fundingState)

	// Check if funding limit is exceeded
	if len(activeFundings) >= types.MaxFunders {
		lowestFunding, err := k.GetLowestFunding(activeFundings)
		if err != nil {
			util.PanicHalt(k.upgradeKeeper, ctx, err.Error())
		}

		// Check if lowest funding is lower than new funding
		// TODO: what criteria should we use to determine if a funder is the lowest?
		if lowestFunding.Amount < funding.Amount {
			// If so, check if lowest funding is from someone else
			if lowestFunding.FunderAddress != funding.FunderAddress {
				// Prepare to defund lowest funding
				defunding = lowestFunding
			}
		} else {
			return nil, errors.Wrapf(errorsTypes.ErrLogic, types.ErrFundsTooLow.Error(), lowestFunding.Amount)
		}
	}

	// User is allowed to fund
	// Let's see if he has enough funds
	if err := util.TransferFromAddressToModule(k.bankKeeper, ctx, msg.Creator, types.ModuleName, msg.Amount); err != nil {
		//if err := util.TransferFromAddressToModule(k.bankKeeper, ctx, msg.Creator, "pool", msg.Amount); err != nil {
		return nil, err
	}

	// Check if defunding is necessary
	if defunding != nil {
		err := k.defundLowestFunding(ctx, defunding, &fundingState, msg.PoolId)

		// TODO: what to do if defunding fails? Should we return the funds to the user?
		if err2 := util.TransferFromModuleToAddress(k.bankKeeper, ctx, types.ModuleName, msg.Creator, msg.Amount); err2 != nil {
			k.Logger(ctx).Error("Failed to defund lowest funding", "error", err.Error())
			k.Logger(ctx).Error("Failed to transfer funds back to user", "error", err2.Error())
			return nil, err2
		}
		if err != nil {
			return nil, err
		}
	}

	// Funding must be active
	fundingState.SetActive(&funding)

	// Save funding and funding state
	k.SetFunding(ctx, &funding)
	k.SetFundingState(ctx, &fundingState)

	// Emit a fund event.
	_ = ctx.EventManager().EmitTypedEvent(&types.EventFundPool{
		PoolId:          msg.PoolId,
		Address:         msg.Creator,
		Amount:          msg.Amount,
		AmountPerBundle: msg.AmountPerBundle,
	})

	return &types.MsgFundPoolResponse{}, nil
}
