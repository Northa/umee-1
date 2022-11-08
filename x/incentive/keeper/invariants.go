package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/umee-network/umee/v3/x/incentive/types"
)

const (
	routeNextID = "next-id"
)

// RegisterInvariants registers the incentive module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, routeNextID, NextIDInvariant(k))
}

// AllInvariants runs all invariants of the x/incentive module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		return NextIDInvariant(k)(ctx)
	}
}

// NextIDInvariant checks that next ID is nonzero
func NextIDInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		// TODO: create invariants

		var count int

		// TODO: check that next ID = len(get all programs) + 1
		// and that all program IDs are < next ID
		// and that no program IDs repeat

		// TODO: consider moving this to validate genesis

		broken := count != 0

		return sdk.FormatInvariant(
			types.ModuleName, routeNextID,
			fmt.Sprintf("invalid next ID: %d", count),
		), broken
	}
}
