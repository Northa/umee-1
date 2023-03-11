package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/umee-network/umee/v4/x/leverage/types"
)

// inspectorFilter defines a function which decides whether to pay attention to a BorrowerSummary
type inspectorFilter func(*types.BorrowerSummary) bool

func withMinBorrowedValue(value sdk.Dec) inspectorFilter {
	return func(bs *types.BorrowerSummary) bool {
		return bs.BorrowedValue.GTE(value)
	}
}

func withMinCollateralValue(value sdk.Dec) inspectorFilter {
	return func(bs *types.BorrowerSummary) bool {
		return bs.CollateralValue.GTE(value)
	}
}

func withMinSuppliedlValue(value sdk.Dec) inspectorFilter {
	return func(bs *types.BorrowerSummary) bool {
		return bs.SuppliedValue.GTE(value)
	}
}

func withMinThresholdUsage(value sdk.Dec) inspectorFilter {
	return func(bs *types.BorrowerSummary) bool {
		return bs.LiquidationThreshold.IsPositive() && bs.BorrowedValue.Quo(bs.LiquidationThreshold).GTE(value)
	}
}

func withMinLTV(value sdk.Dec) inspectorFilter {
	return func(bs *types.BorrowerSummary) bool {
		return bs.CollateralValue.IsPositive() && bs.BorrowedValue.Quo(bs.CollateralValue).GTE(value)
	}
}

func withMinLimitUsage(value sdk.Dec) inspectorFilter {
	return func(bs *types.BorrowerSummary) bool {
		return bs.BorrowLimit.IsPositive() && bs.BorrowedValue.Quo(bs.BorrowLimit).GTE(value)
	}
}

// withZeroes returns borrower summaries with unexpected zero values (knowing that borrower summaries only exist
// for accounts with borrowed tokens)
func withZeroes(value sdk.Dec) inspectorFilter {
	return func(bs *types.BorrowerSummary) bool {
		return bs.CollateralValue.IsZero() || bs.LiquidationThreshold.IsZero() || bs.BorrowedValue.IsZero()
	}
}
