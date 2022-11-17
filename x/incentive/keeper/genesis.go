package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/umee-network/umee/v3/x/incentive/types"
)

// InitGenesis initializes the x/incentive module state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	if err := k.setParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	// TODO: Set everything else
}

// ExportGenesis returns the x/incentive module's exported genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		// TODO: Get everything
		k.GetParams(ctx),
		nil,
		0,
		sdk.NewCoins(),
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}
