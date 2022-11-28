package incentive

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	gogotypes "github.com/gogo/protobuf/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "incentive"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// StoreKey defines the query route
	QuerierRoute = ModuleName
)

// KVStore key prefixes
var (
	// Individually store params from MsgGovSetParams
	KeyPrefixParamMaxUnbondings           = []byte{0x01, 0x01}
	KeyPrefixParamUnbondingDurationLong   = []byte{0x01, 0x02}
	KeyPrefixParamUnbondingDurationMiddle = []byte{0x01, 0x03}
	KeyPrefixParamUnbondingDurationShort  = []byte{0x01, 0x04}
	KeyPrefixParamTierWeightShort         = []byte{0x01, 0x05}
	KeyPrefixParamTierWeightMiddle        = []byte{0x01, 0x06}

	// Regular state
	KeyPrefixUpcomingIncentiveProgram  = []byte{0x02}
	KeyPrefixOngoingIncentiveProgram   = []byte{0x03}
	KeyPrefixCompletedIncentiveProgram = []byte{0x04}
	KeyPrefixNextProgramID             = []byte{0x05}
	KeyPrefixLastRewardsTime           = []byte{0x06}
	KeyPrefixTotalBonded               = []byte{0x07}
	KeyPrefixBondAmount                = []byte{0x08}
	KeyPrefixPendingReward             = []byte{0x09}
	KeyPrefixRewardBasis               = []byte{0x0A}
	KeyPrefixRewardAccumulator         = []byte{0x0B}
	KeyPrefixUnbonding                 = []byte{0x0C}
)

// CreateUpcomingIncentiveProgramKey returns a KVStore key for getting and setting an upcoming IncentiveProgram.
func CreateUpcomingIncentiveProgramKey(cdc codec.Codec, id uint32) []byte {
	// prefix | id
	var key []byte
	key = append(key, KeyPrefixUpcomingIncentiveProgram...)

	// note: use of codec required by using a uint32 as part of a key
	bz, err := cdc.Marshal(&gogotypes.UInt32Value{Value: id})
	if err != nil {
		panic(err)
	}

	key = append(key, bz...)
	return key
}

// CreateOngoingIncentiveProgramKey returns a KVStore key for getting and setting an ongoing IncentiveProgram.
func CreateOngoingIncentiveProgramKey(cdc codec.Codec, id uint32) []byte {
	// prefix | id
	var key []byte
	key = append(key, KeyPrefixOngoingIncentiveProgram...)

	// note: use of codec required by using a uint32 as part of a key
	bz, err := cdc.Marshal(&gogotypes.UInt32Value{Value: id})
	if err != nil {
		panic(err)
	}

	key = append(key, bz...)
	return key
}

// CreateCompletedIncentiveProgramKey returns a KVStore key for getting and setting an completed IncentiveProgram.
func CreateCompletedIncentiveProgramKey(cdc codec.Codec, id uint32) []byte {
	// prefix | id
	var key []byte
	key = append(key, KeyPrefixCompletedIncentiveProgram...)

	// note: use of codec required by using a uint32 as part of a key
	bz, err := cdc.Marshal(&gogotypes.UInt32Value{Value: id})
	if err != nil {
		panic(err)
	}

	key = append(key, bz...)
	return key
}

// CreateTotalBondedKey returns a KVStore key for getting and setting the
// total bonded amount tracker for a single uToken.
func CreateTotalBondedKey(uTokenDenom string) []byte {
	// prefix | denom | 0x00
	var key []byte
	key = append(key, KeyPrefixTotalBonded...)
	key = append(key, []byte(uTokenDenom)...)
	return append(key, 0) // append 0 for null-termination
}

// CreateBondedAmountKey returns a KVStore key for getting and setting a
// bonded amount for a denom and address.
func CreateBondedAmountKey(addr sdk.AccAddress, uTokenDenom string) []byte {
	// prefix | lengthprefixed(addr) | denom | 0x00
	key := CreateBondedAmountKeyNoDenom(addr)
	key = append(key, []byte(uTokenDenom)...)
	return append(key, 0) // append 0 for null-termination
}

// CreateBondedAmountKeyNoDenom returns the common prefix used by all bonded tokens
// associated with a given address.
func CreateBondedAmountKeyNoDenom(addr sdk.AccAddress) []byte {
	// prefix | lengthprefixed(addr)
	var key []byte
	key = append(key, KeyPrefixBondAmount...)
	key = append(key, address.MustLengthPrefix(addr)...)
	return key
}

// CreatePendingRewardKeyNoDenom returns a KVStore key for getting all denoms
// of pending rewards assoicated with an account.
func CreatePendingRewardKeyNoDenom(addr sdk.AccAddress) []byte {
	// prefix | lengthprefixed(addr)
	var key []byte
	key = append(key, KeyPrefixPendingReward...)
	key = append(key, address.MustLengthPrefix(addr)...)
	return key
}

// CreatePendingRewardKey returns a KVStore key for getting and setting the
// amount of rewards for a given address and reward denom which have been
// calculated but not yet claimed.
func CreatePendingRewardKey(addr sdk.AccAddress, denom string) []byte {
	// prefix | lengthprefixed(addr) | denom | 0x00
	key := CreatePendingRewardKeyNoDenom(addr)
	key = append(key, []byte(denom)...)
	return append(key, 0) // append 0 for null-termination
}

// CreateRewardBasisKey returns a KVStore key for getting and setting the
// reward basis in a single reward denom for a given bonded uToken denom and address.
func CreateRewardBasisKey(addr sdk.AccAddress, lockDenom, rewardDenom string) []byte {
	// prefix | lengthprefixed(addr) | bondDenom | 0x00 | rewardDenom | 0x00
	key := CreateRewardBasisKeyNoRewardDenom(addr, lockDenom)
	key = append(key, []byte(rewardDenom)...)
	return append(key, 0) // append 0 for null-termination
}

// CreateRewardBasisKeyNoRewardDenom returns a KVStore key for getting and setting the
// reward bases in all reward denoms for a given bonded uToken denom and address.
func CreateRewardBasisKeyNoRewardDenom(addr sdk.AccAddress, lockDenom string) []byte {
	// prefix | lengthprefixed(addr) | denom | 0x00
	key := CreateRewardBasisKeyNoBondDenom(addr)
	key = append(key, []byte(lockDenom)...)
	return append(key, 0) // append 0 for null-termination
}

// CreateRewardBasisKeyNoBondDenom returns the common prefix used by all reward bases
// associated with a given address.
func CreateRewardBasisKeyNoBondDenom(addr sdk.AccAddress) []byte {
	// prefix | lengthprefixed(addr)
	var key []byte
	key = append(key, KeyPrefixRewardBasis...)
	key = append(key, address.MustLengthPrefix(addr)...)
	return key
}

// CreateRewardAccumulatorKey returns a KVStore key for getting and setting the
// reward basis tracker for given bonded uToken and reward token denoms.
func CreateRewardAccumulatorKey(bondDenom, rewardDenom string) []byte {
	// prefix | bondDenom | 0x00 | rewardDenom | 0x00
	key := CreateRewardAccumulatorKeyNoRewardDenom(bondDenom)
	key = append(key, []byte(rewardDenom)...)
	return append(key, 0) // append 0 for null-termination
}

// CreateRewardAccumulatorKeyNoRewardDenom returns a KVStore key for getting and setting the
// reward basis tracker for all reward denoms associated with a single bonded uToken denom.
func CreateRewardAccumulatorKeyNoRewardDenom(bondDenom string) []byte {
	// prefix | bondDenom | 0x00
	var key []byte
	key = append(key, KeyPrefixRewardAccumulator...)
	key = append(key, []byte(bondDenom)...)
	return append(key, 0) // append 0 for null-termination
}

// CreateUnbondingKey returns a KVStore key for storing an unbonding associated with a given
// account and height.
func CreateUnbondingKey(cdc codec.Codec, addr sdk.AccAddress, height uint64) []byte {
	// prefix | lengthprefixed(addr) | height
	var key []byte
	key = append(key, KeyPrefixUnbonding...)
	key = append(key, address.MustLengthPrefix(addr)...)

	// note: use of codec required by using a uint64 as part of a key
	bz, err := cdc.Marshal(&gogotypes.UInt64Value{Value: height})
	if err != nil {
		panic(err)
	}
	key = append(key, bz...)

	return key
}
