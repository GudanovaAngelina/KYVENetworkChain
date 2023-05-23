package keeper

import (
	"fmt"

	"github.com/KYVENetwork/chain/util"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// Delegation
	delegationKeeper "github.com/KYVENetwork/chain/x/delegation/keeper"
	// Stakers
	"github.com/KYVENetwork/chain/x/stakers/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storeTypes.StoreKey
		memKey   storeTypes.StoreKey

		authority string

		accountKeeper util.AccountKeeper
		bankKeeper    util.BankKeeper
		distrkeeper   util.DistributionKeeper
		poolKeeper    types.PoolKeeper
		upgradeKeeper util.UpgradeKeeper
		// TODO(@john): Switch this back once app wiring works.
		delegationKeeper delegationKeeper.Keeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storeTypes.StoreKey,
	memKey storeTypes.StoreKey,

	authority string,

	accountKeeper util.AccountKeeper,
	bankKeeper util.BankKeeper,
	distrkeeper util.DistributionKeeper,
	poolKeeper types.PoolKeeper,
	upgradeKeeper util.UpgradeKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		authority: authority,

		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		distrkeeper:   distrkeeper,
		poolKeeper:    poolKeeper,
		upgradeKeeper: upgradeKeeper,
	}
}

func (k *Keeper) SetDelegationKeeper(delegationKeeper delegationKeeper.Keeper) {
	k.delegationKeeper = delegationKeeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) StoreKey() storeTypes.StoreKey {
	return k.storeKey
}
