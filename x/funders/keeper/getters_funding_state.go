package keeper

import (
	"errors"
	"fmt"
	"github.com/KYVENetwork/chain/x/funders/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DoesFundingStateExist checks if the FundingState exists
func (k Keeper) DoesFundingStateExist(ctx sdk.Context, poolId uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.FundingStateKeyPrefix)
	return store.Has(types.FundingStateKey(poolId))
}

// GetFundingState returns the FundingState
func (k Keeper) GetFundingState(ctx sdk.Context, poolId uint64) (fundingState types.FundingState, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.FundingStateKeyPrefix)

	b := store.Get(types.FundingStateKey(
		poolId,
	))
	if b == nil {
		return fundingState, false
	}

	k.cdc.MustUnmarshal(b, &fundingState)
	return fundingState, true
}

// SetFundingState sets a specific FundingState in the store from its index
func (k Keeper) setFundingState(ctx sdk.Context, fundingState types.FundingState) {
	b := k.cdc.MustMarshal(&fundingState)
	storeByFunder := prefix.NewStore(ctx.KVStore(k.storeKey), types.FundingStateKeyPrefix)
	storeByFunder.Set(types.FundingStateKey(
		fundingState.PoolId,
	), b)
}

func (k Keeper) GetActiveFundings(ctx sdk.Context, fundingState types.FundingState) (fundings []types.Funding) {
	for _, funder := range fundingState.ActiveFunderAddresses {
		funding, found := k.GetFunding(ctx, funder, fundingState.PoolId)
		if found {
			fundings = append(fundings, funding)
		} // else should never happen or we have a corrupted state
	}
	return fundings
}

// GetLowestFunding returns the funding with the lowest amount
// Precondition: len(fundings) > 0
func (k Keeper) GetLowestFunding(fundings []types.Funding) (lowestFunding *types.Funding, err error) {
	if len(fundings) == 0 {
		return nil, errors.New(fmt.Sprintf("no active fundings"))
	}

	for _, funding := range fundings {
		if funding.Amount < lowestFunding.Amount {
			lowestFunding = &funding
		}
	}
	return lowestFunding, nil
}