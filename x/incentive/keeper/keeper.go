package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/umee-network/umee/v3/x/incentive"
)

type Keeper struct {
	cdc            codec.Codec
	storeKey       storetypes.StoreKey
	bankKeeper     incentive.BankKeeper
	leverageKeeper incentive.LeverageKeeper
}

func NewKeeper(
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
	bk incentive.BankKeeper,
	lk incentive.LeverageKeeper,
) Keeper {
	return Keeper{
		cdc:            cdc,
		storeKey:       storeKey,
		bankKeeper:     bk,
		leverageKeeper: lk,
	}
}

func (k Keeper) kvStore(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", incentive.ModuleName))
}

// ModuleBalance returns the amount of a given token held in the x/incentive module account
func (k Keeper) ModuleBalance(ctx sdk.Context, denom string) sdk.Coin {
	amount := k.bankKeeper.SpendableCoins(ctx, authtypes.NewModuleAddress(incentive.ModuleName)).AmountOf(denom)
	return sdk.NewCoin(denom, amount)
}

// Claim claims any pending rewards belonging to an address.
func (k Keeper) Claim(ctx sdk.Context, addr sdk.AccAddress) (sdk.Coins, error) {
	// calculate and set pending rewards available to account at the current block
	if err := k.AllocatePendingRewards(ctx, addr); err != nil {
		return sdk.Coins{}, err
	}

	// get the sum of newly calculated and previously pending rewards
	rewards := k.GetPendingRewards(ctx, addr)

	// claim all pending rewards
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, incentive.ModuleName, addr, rewards); err != nil {
		return sdk.Coins{}, err
	}

	// set pending rewards to zero
	if err := k.SetPendingRewards(ctx, addr, sdk.NewCoins()); err != nil {
		return sdk.Coins{}, err
	}

	// returns the amount received
	return rewards, nil
}
